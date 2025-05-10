// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/dto"
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IAuth interface {
		AuthorizationToken(ctx context.Context, userUUID string) (*dto.AuthorizeTokenDTO, *berror.ErrorCode)
		// CheckUserExistByUsername 检查指定用户名的用户是否存在。
		//
		// 参数:
		//   - ctx: 请求上下文信息，用于控制操作的生命周期。
		//   - Username: 需要检查的用户名。
		//
		// 返回:
		//   - err: 错误代码，表示用户不存在或其他错误情况。
		CheckUserExistByUsername(ctx context.Context, Username string) (err *berror.ErrorCode)
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
		UserLogin(ctx context.Context, Username string, Password string) (*dto.UserInfoDTO, *berror.ErrorCode)
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
