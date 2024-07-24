/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package auth

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

// CheckUserHasSuperAdmin
//
// # 检查用户是否有超级管理员权限
//
// 检查用户是否有超级管理员权限，用于检查用户是否有超级管理员权限；
// 用于检查用户是否有超级管理员权限；
//
// # 参数
//   - ctx				上下文(context.Context)
//   - authorization	用户唯一令牌(string)
//
// # 返回
//   - error	错误信息
func (s *sAuth) CheckUserHasSuperAdmin(ctx context.Context, authorization string) (err error) {
	g.Log().Notice(ctx, "[SERV] auth.CheckUserHasSuperAdmin | 检查用户是否有超级管理员权限")
	// 通过用户唯一令牌查询用户信息
	getUserUUID, err := g.Redis().Get(ctx, "token:"+butil.TokenRemoveBearer(authorization))
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "用户令牌无效")
	}
	// 查询用户信息
	var getSuperAdminUUID *entity.Info
	err = dao.Info.Ctx(ctx).Where(do.Info{Key: "system_admin"}).Scan(&getSuperAdminUUID)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户信息失败")
	}
	if strings.EqualFold(getUserUUID.String(), getSuperAdminUUID.Value) {
		return berror.NewError(bcode.ServerInternalError, "用户无权限")
	}
	return nil
}

// CheckUserHasAdmin
//
// # 检查用户是否有管理员权限
//
// 检查用户是否有管理员权限，用于检查用户是否有管理员权限；
// 用于检查用户是否有管理员权限；
//
// # 参数
//   - ctx				上下文(context.Context)
//   - authorization	用户唯一令牌(string)
//
// # 返回
//   - error	错误信息
func (s *sAuth) CheckUserHasAdmin(ctx context.Context, authorization string) (err error) {
	g.Log().Notice(ctx, "[SERV] auth.CheckUserHasAdmin | 检查用户是否有管理员权限")
	getUser, err := s.GetUserByAuthorization(ctx, authorization)
	if err != nil {
		return err
	}
	// 获取所在角色
	var getRole *entity.Role
	err = dao.Role.Ctx(ctx).Where(do.Role{Ruuid: getUser.Ruuid}).Scan(&getRole)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户角色失败")
	}
	if getRole.Name == "admin" {
		return nil
	} else {
		return berror.NewError(bcode.Unauthorized, "用户无权限")
	}
}

// CheckUserHasLogin
//
// # 检查用户是否登录
//
// 检查用户是否登录，用于检查用户是否登录；
//
// # 参数
//   - ctx				上下文(context.Context)
//   - authorization	用户唯一令牌(string)
//
// # 返回
//   - error	错误信息
func (s *sAuth) CheckUserHasLogin(ctx context.Context, authorization string) (err error) {
	g.Log().Notice(ctx, "[SERV] auth.CheckUserHasLogin | 检查用户是否登录")
	getUser, err := s.GetUserByAuthorization(ctx, authorization)
	if err != nil {
		return err
	}
	// 获取用户所在角色组为非禁用状态
	var getRole *entity.Role
	err = dao.Role.Ctx(ctx).Where(do.Role{Ruuid: getUser.Ruuid}).Scan(&getRole)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户角色失败")
	}
	if getRole.Name == "bad" {
		return berror.NewError(bcode.Unauthorized, "用户无权限")
	} else {
		return nil
	}
}

// GetUserByAuthorization
//
// # 通过用户唯一令牌获取用户信息
//
// 通过用户唯一令牌获取用户信息，用于通过用户唯一令牌获取用户信息；
// 用于通过用户唯一令牌获取用户信息；
//
// # 参数
//   - ctx				上下文(context.Context)
//   - authorization	用户唯一令牌(string)
//
// # 返回
//   - user		用户信息(*entity.User)
//   - error	错误信息
func (s *sAuth) GetUserByAuthorization(ctx context.Context, authorization string) (user *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] auth.GetUserByAuthorization | 通过用户唯一令牌获取用户信息")
	// 通过用户唯一令牌查询用户信息
	getUserUUID, err := g.Redis().Get(ctx, "token:"+butil.TokenRemoveBearer(authorization))
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "用户令牌无效")
	}
	// 查询用户信息
	err = dao.User.Ctx(ctx).Where(do.User{Uuid: getUserUUID.String()}).Scan(&user)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户信息失败")
	}
	if user == nil {
		return nil, berror.NewError(bcode.NotExist, "用户不存在")
	}
	return user, nil
}
