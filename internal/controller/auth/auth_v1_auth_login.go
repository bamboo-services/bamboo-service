package auth

import (
	"bamboo-service/api/auth/v1"
	"bamboo-service/internal/model/dto/dsingle"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
)

func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	iAuth := service.Auth()
	// 登录
	userInfo, errorCode := iAuth.UserLogin(ctx, req.Username, req.Password)
	if errorCode != nil {
		blog.ControllerError(ctx, "AuthLogin", errorCode.Error())
		return nil, errorCode
	}
	// 授权
	getToken, errorCode := iAuth.AuthorizationToken(ctx, userInfo.UserUuid)
	if errorCode != nil {
		blog.ControllerError(ctx, "AuthLogin", errorCode.Error())
		return nil, errorCode
	}
	return &v1.AuthLoginRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "登录成功", &dsingle.UserLoginDTO{
			UserInfo: userInfo,
			Token:    getToken,
		}),
	}, nil
}
