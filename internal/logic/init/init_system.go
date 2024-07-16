/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package init

import (
	"bamboo-service/internal/constant"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
)

// SetSystemWeb
//
// # 初始化系统网站
//
// 该接口用于确认前端地址，用于系统初始化；
// 否则系统将不放行 CORS 请求；
// 该接口需要在系统初始化时调用；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - referer		来源地址(string)
//
// # 返回
//   - err		错误信息(error)
func (s *sInit) SetSystemWeb(ctx context.Context, referer string) (err error) {
	g.Log().Notice(ctx, "[LOGIC] init.SetSystemWeb | 初始化系统网站")
	if !constant.InitializeMode {
		return berror.NewError(bcode.ServerInternalError, "系统已经初始化，无法再次初始化系统网站")
	}
	// 设置系统网站
	constant.SystemReferer = referer
	// 将配置信息更新到数据库
	_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "system_referer"}).Update(do.Info{Value: referer})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "更新系统网站地址失败")
	}
	return nil
}

// SetSystemAdmin
//
// # 初始化系统管理员
//
// 该接口用于初始化系统管理员，用于系统初始化；
// 该接口需要在系统初始化时调用；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - username		用户名(string)
//   - phone		手机号(string)
//   - email		邮箱(string)
//   - password	密码(string)
//
// # 返回
//   - err		错误信息(error)
func (s *sInit) SetSystemAdmin(ctx context.Context, username, phone, email, password string) (err error) {
	g.Log().Notice(ctx, "[LOGIC] init.SetSystemAdmin | 初始化系统管理员")
	if !constant.InitializeMode {
		return berror.NewError(bcode.ServerInternalError, "系统已经初始化，无法再次初始化系统管理员")
	}
	// 获取系统管理员
	var getAdmin *entity.Info
	err = dao.Info.Ctx(ctx).Where(do.Info{Key: "system_admin"}).Scan(&getAdmin)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "获取系统管理员失败")
	}
	// 设置系统管理员
	_, err = dao.User.Ctx(ctx).Where(do.User{Uuid: getAdmin.Value}).Update(do.User{
		Username: username,
		Phone:    phone,
		Email:    email,
		Password: butil.PasswordEncode(password),
	})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "更新系统管理员失败")
	}
	// 关闭初始化模式
	constant.InitializeMode = false
	// 将配置信息更新到数据库
	_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "has_initial_mode"}).Update(do.Info{Value: "0"})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "关闭初始化模式失败")
	}
	return nil
}
