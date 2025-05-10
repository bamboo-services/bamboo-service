package auth

import (
	"bamboo-service/api/auth/v1"
	"bamboo-service/internal/model/dto/dsingle"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
)

// AuthLogin 用户登录并获取授权令牌。
//
// 参数:
//   - ctx: 上下文信息。
//   - req: 用户登录请求，包含用户名和密码。
//
// 返回:
//   - res: 用户登录响应，包含用户基本信息和授权令牌。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	blog.ControllerInfo(ctx, "AuthLogin", "用户登录")
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
