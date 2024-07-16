/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package user

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

// UserChangePassword
//
// # 用户修改密码
//
// 用户修改密码，用于用户修改密码，需要传递用户的 UUID 和新密码；
// 该接口将会对用户的密码进行修改，修改成功后将会返回 nil 信息；
// 新的密码将会重新加密存入数据库中，旧密码将会保存在数据库旧密码位置；
// 修改密码将会检查不允许修改当前密码以及上一次密码，以及重置密码阶段可以对上一次的密码进行密码找回操作；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - getUUID		用户 UUID(uuid.UUID)
//   - newPassword	新密码(string)
//
// # 返回
//   - err			错误信息(error)
func (s *sUser) UserChangePassword(ctx context.Context, getUUID uuid.UUID, newPassword string) (err error) {
	g.Log().Notice(ctx, "[SERV] user.UserChangePassword | 用户修改密码接口")
	// 通过 UUID 对用户数据进行获取
	var getUser *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{Uuid: getUUID}).Scan(&getUser)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户失败")
	}
	if getUser == nil {
		return berror.NewError(bcode.NotExist, "用户不存在")
	}
	// 进行密码校验
	if butil.PasswordVerify(newPassword, getUser.Password) {
		return berror.NewError(bcode.OperationFailed, "修改的密码与当前密码相同")
	}
	// 检查是否存在旧密码
	if getUser.OldPassword == "" {
		if butil.PasswordVerify(newPassword, getUser.OldPassword) {
			return berror.NewError(bcode.OperationFailed, "修改的密码与旧密码相同")
		}
	}
	// 对密码进行替换后加密
	getUser.OldPassword = getUser.Password
	getUser.Password = butil.PasswordEncode(newPassword)
	// 更新用户数据
	_, err = dao.User.Ctx(ctx).Where(do.User{Uuid: getUUID}).Update(getUser)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "更新用户失败")
	}
	return nil
}
