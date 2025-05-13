package user

import (
	"bamboo-service/internal/custom"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
)

// CheckUserEmailIsVerify 检查指定邮箱是否已验证。
//
// 参数:
//   - ctx: 请求上下文信息，用于控制操作的生命周期。
//   - email: 需要检查的邮箱地址。
//
// 返回:
//   - *berror.ErrorCode: 错误代码，表示邮箱未验证或其他错误情况。返回 nil 表示邮箱已验证。
func (s *sUser) CheckUserEmailIsVerify(ctx context.Context, email string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckUserEmailIsVerify", "检查邮箱 %s 是否已验证", email)
	getUser, errorCode := s.GetUserByEmail(ctx, email)
	if errorCode != nil {
		return errorCode
	}
	if getUser.EmailVerifiedAt == nil {
		return custom.ErrorEmailNotVerify
	}
	return nil
}

// CheckUserPhoneIsVerify 检查指定手机号是否已验证。
//
// 参数:
//   - ctx: 请求上下文信息，用于控制操作的生命周期。
//   - phone: 需要检查的手机号。
//
// 返回:
//   - *berror.ErrorCode: 错误代码，表示手机号未验证或其他错误情况。返回 nil 表示手机号已验证。
func (s *sUser) CheckUserPhoneIsVerify(ctx context.Context, phone string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckUserPhoneIsVerify", "检查手机号 %s 是否已验证", phone)
	getUser, errorCode := s.GetUserByPhone(ctx, phone)
	if errorCode != nil {
		return errorCode
	}
	if getUser.PhoneVerifiedAt == nil {
		return custom.ErrorEmailNotVerify
	}
	return nil
}
