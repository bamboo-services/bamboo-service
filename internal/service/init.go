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
	IInit interface {
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
		SetSystemWeb(ctx context.Context, referer string) (err error)
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
		SetSystemAdmin(ctx context.Context, username, phone, email, password string) (err error)
	}
)

var (
	localInit IInit
)

func Init() IInit {
	if localInit == nil {
		panic("implement not found for interface IInit, forgot register?")
	}
	return localInit
}

func RegisterInit(i IInit) {
	localInit = i
}
