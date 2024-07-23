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
