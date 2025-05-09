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
		// UserLogin 处理用户登录请求，通过用户名和密码验证用户身份，返回错误信息（如果有）。
		UserLogin(ctx context.Context, Username string, Password string) (userInfo *dto.UserInfoDTO, err *berror.ErrorCode)
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
