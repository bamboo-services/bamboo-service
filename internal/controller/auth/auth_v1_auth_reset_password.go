package auth

import (
	"bamboo-service/internal/custom"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"

	"bamboo-service/api/auth/v1"
)

// AuthResetPassword 处理用户重置密码请求。
//
// 参数:
//   - ctx: 上下文信息。
//   - req: 密码重置请求，包含用户邮箱、验证码和新密码。
//
// 返回:
//   - res: 密码重置响应，表示重置结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) AuthResetPassword(ctx context.Context, req *v1.AuthResetPasswordReq) (res *v1.AuthResetPasswordRes, err error) {
	blog.ControllerInfo(ctx, "AuthResetPassword", "重置密码")
	// 密码一致性验证
	if req.Password != req.ConfirmPassword {
		return nil, custom.ErrorUserConfirmPasswordIncorrect
	}

	// 检查邮箱是否存在
	iUser := service.User()
	errorCode := iUser.CheckUserExistByEmail(ctx, req.Email)
	if errorCode != nil {
		return nil, errorCode
	}

	// 检查该用户邮件是否已验证
	errorCode = iUser.CheckUserEmailIsVerify(ctx, req.Email)
	if errorCode != nil {
		return nil, errorCode
	}

	// 检查验证码是否正确
	iMail := service.Mail()
	errorCode = iMail.VerifyMailCode(ctx, "reset_password", req.Email, req.Code)
	if errorCode != nil {
		return nil, errorCode
	}

	// 重置密码
	iAuth := service.Auth()
	errorCode = iAuth.ResetPassword(ctx, req.Email, req.Password)
	if errorCode != nil {
		return nil, errorCode
	}

	return &v1.AuthResetPasswordRes{
		ResponseDTO: bresult.Success(ctx, "密码重置成功"),
	}, nil
}
