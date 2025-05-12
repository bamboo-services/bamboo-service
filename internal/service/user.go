// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IUser interface {
		// CheckUserExistByUsername 检查指定用户名的用户是否存在。
		//
		// 当用户
		//
		// 参数:
		//   - ctx: 请求上下文信息，用于控制操作的生命周期。
		//   - Username: 需要检查的用户名。
		//
		// 返回:
		//   - err: 错误代码，表示用户不存在或其他错误情况。
		CheckUserExistByUsername(ctx context.Context, username string) *berror.ErrorCode
		// CheckUserExistByEmail 检查指定邮箱的用户是否存在。
		//
		// 参数:
		//   - ctx: 请求上下文，用于控制操作生命周期。
		//   - Email: 需要检查的邮箱地址。
		//
		// 返回:
		//   - 错误代码，表示用户是否存在或其他错误情况。
		CheckUserExistByEmail(ctx context.Context, email string) *berror.ErrorCode
		// CheckUserExistByPhone 检查指定手机号的用户是否存在。
		//
		// 参数:
		//   - ctx: 请求上下文，用于控制操作生命周期。
		//   - phone: 需要检查的手机号。
		//
		// 返回:
		//   - 错误代码，表示用户是否存在或其他错误情况。
		CheckUserExistByPhone(ctx context.Context, phone string) *berror.ErrorCode
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
