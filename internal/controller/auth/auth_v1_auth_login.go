package auth

import (
	"bamboo-service/api/auth/v1"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
)

func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	iAuth := service.Auth()
	userInfo, errorCode := iAuth.UserLogin(ctx, req.Username, req.Password)
	if errorCode != nil {
		blog.ControllerDebug(ctx, "AuthLogin", errorCode.Error())
		return nil, errorCode
	}
	return &v1.AuthLoginRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "登录成功", userInfo),
	}, nil
}
