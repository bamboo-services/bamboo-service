package auth

import (
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"

	"bamboo-service/api/auth/v1"
)

func (c *ControllerV1) AuthResetPassword(ctx context.Context, req *v1.AuthResetPasswordReq) (res *v1.AuthResetPasswordRes, err error) {
	blog.ControllerInfo(ctx, "AuthResetPassword", "重置密码")

	// 检查邮箱是否存在
	iUser := service.User()
	errorCode := iUser.CheckUserExistByEmail(ctx, req.Email)
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
	// TODO-25051301 密码重置逻辑

	return &v1.AuthResetPasswordRes{
		ResponseDTO: bresult.Success(ctx, "密码重置成功"),
	}, nil
}
