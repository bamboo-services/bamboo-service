/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/entity"
	"context"
)

type (
	IAuth interface {
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
		CheckUserHasSuperAdmin(ctx context.Context, authorization string) (err error)
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
		CheckUserHasAdmin(ctx context.Context, authorization string) (err error)
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
		CheckUserHasLogin(ctx context.Context, authorization string) (err error)
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
		GetUserByAuthorization(ctx context.Context, authorization string) (user *entity.User, err error)
		// IsUserCanDo
		//
		// # 检查用户是否有权限
		//
		// 检查用户是否有权限，用于检查用户是否有权限；
		// 用于服务层之间进行用户是否有权限操作；
		// 一般用作个人业务，个人有权限进行操作，有权限的人也可以进行操作，超级管理员也可以进行操作；
		//
		// # 参数
		//   - ctx				上下文(context.Context)
		//   - authorization	用户唯一令牌(string)
		//   - userUUID			用户唯一标识(string)
		//   - permission		权限(string)
		//
		// # 返回
		//   - error	错误信息
		IsUserCanDo(ctx context.Context, authorization, userUUID string, permission string) (err error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
