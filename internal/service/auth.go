// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "bamboo-service/api/auth/v1"
	"bamboo-service/internal/model/dto"
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IAuth interface {
		// AuthorizationToken 生成用户的授权令牌并存储到 Redis 缓存中。
		//
		// 参数:
		//   - ctx: 请求上下文信息。
		//   - userUUID: 用户唯一标识符。
		//
		// 返回:
		//   - *dto.AuthorizeTokenDTO: 包含生成的授权令牌及其相关信息。
		//   - *berror.ErrorCode: 错误代码，表示可能的存储或其他错误情况。
		AuthorizationToken(ctx context.Context, userUUID string) (*dto.AuthorizeTokenDTO, *berror.ErrorCode)
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
		// UserLogin 验证用户登录，并返回用户信息。
		//
		// 参数:
		//   - ctx: 请求上下文。
		//   - Username: 用户名。
		//   - Password: 用户密码。
		//
		// 返回:
		//   - userInfo: 用户信息数据传输对象。
		//   - err: 错误代码，表示登录失败的原因。
		UserLogin(ctx context.Context, username string, password string) (*dto.UserInfoDTO, *berror.ErrorCode)
		// UserRegister 注册新用户。
		//
		// 参数:
		//   - ctx: 请求上下文，用于控制操作生命周期。
		//   - request: 包含用户名和密码的用户注册请求。
		//
		// 返回:
		//   - 错误代码，表示注册失败的原因或 nil 表示成功。
		UserRegister(ctx context.Context, request *v1.AuthLoginReq) berror.ErrorCode
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
