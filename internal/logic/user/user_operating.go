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
)

// UserRegister
//
// # 用户注册
//
// 用户注册，用于用户注册，需要传递用户名、邮箱、手机号、密码；
// 该接口将会对用户的数据进行注册，注册成功后将会返回 nil 信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - username		用户名(string)
//   - email		邮箱(string)
//   - phone		手机号(string)
//   - password		密码(string)
//
// # 返回
//   - err			错误信息(error)
func (s *sUser) UserRegister(ctx context.Context, username, email, phone, password string) (err error) {
	g.Log().Notice(ctx, "[SERV] user.UserRegister | 用户注册接口")
	// 获取默认账户 RUUID
	var getUserRole *entity.Role
	err = dao.Role.Ctx(ctx).Where(do.Role{Name: "user"}).Scan(&getUserRole)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户角色失败")
	}
	// 组织数据
	user := &do.User{
		Uuid:     butil.GenerateRandUUID(),
		Username: username,
		Phone:    phone,
		Email:    email,
		Password: butil.PasswordEncode(password),
		Ruuid:    getUserRole.Ruuid,
		OtpUuid:  butil.GenerateRandUUID(),
	}
	_, err = dao.User.Ctx(ctx).Insert(user)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "用户注册失败")
	}
	return nil
}
