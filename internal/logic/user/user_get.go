package user

import (
	"bamboo-service/internal/custom"
	"bamboo-service/internal/dao"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
)

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
func (s *sUser) CheckUserExistByUsername(ctx context.Context, username string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckUserExistByUsername", "检查用户名 %s 是否存在", username)
	getUser, errorCode := dao.User.GetUserByUsername(ctx, username)
	if errorCode != nil {
		return errorCode
	}
	if getUser != nil {
		return custom.ErrorUserExist
	}
	return nil
}

// CheckUserExistByEmail 检查指定邮箱的用户是否存在。
//
// 参数:
//   - ctx: 请求上下文，用于控制操作生命周期。
//   - Email: 需要检查的邮箱地址。
//
// 返回:
//   - 错误代码，表示用户是否存在或其他错误情况。
func (s *sUser) CheckUserExistByEmail(ctx context.Context, email string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckUserExistByEmail", "检查邮箱 %s 是否存在", email)
	getUser, errorCode := dao.User.GetUserByEmail(ctx, email)
	if errorCode != nil {
		return errorCode
	}
	if getUser != nil {
		return custom.ErrorUserExist
	}
	return nil
}

// CheckUserExistByPhone 检查指定手机号的用户是否存在。
//
// 参数:
//   - ctx: 请求上下文，用于控制操作生命周期。
//   - phone: 需要检查的手机号。
//
// 返回:
//   - 错误代码，表示用户是否存在或其他错误情况。
func (s *sUser) CheckUserExistByPhone(ctx context.Context, phone string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckUserExistByPhone", "检查手机号 %s 是否存在", phone)
	getUser, errorCode := dao.User.GetUserByPhone(ctx, phone)
	if errorCode != nil {
		return errorCode
	}
	if getUser != nil {
		return custom.ErrorUserExist
	}
	return nil
}
