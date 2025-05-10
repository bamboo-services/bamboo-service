package auth

import (
	"bamboo-service/api/auth/v1"
	"bamboo-service/internal/service"
	"context"
)

func (c *ControllerV1) AuthRegister(ctx context.Context, req *v1.AuthRegisterReq) (res *v1.AuthRegisterRes, err error) {
	iAuth := service.Auth()
	errorCode := iAuth.CheckUserExistByUsername(ctx, req.Username)
	if errorCode != nil {
		return nil, errorCode
	}
	errorCode = iAuth.CheckUserExistByEmail(ctx, req.Email)
	if errorCode != nil {
		return nil, errorCode
	}
	errorCode = iAuth.CheckUserExistByPhone(ctx, req.Phone)
	if errorCode != nil {
		return nil, errorCode
	}

	// 创建用户
	errorCode := iAuth.UserRegister(ctx, req)
	if errorCode != nil {
		return nil, errorCode
	}
}
