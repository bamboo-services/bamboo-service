package auth

import (
	"bamboo-service/api/auth/v1"
	"bamboo-service/internal/consts"
	"bamboo-service/internal/custom"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
)

// AuthRegister 处理用户注册请求。
//
// 参数:
//   - ctx: 上下文信息。
//   - req: 用户注册请求，包含用户名、邮箱、手机号、密码等信息。
//
// 返回:
//   - res: 用户注册响应，包含注册成功的用户信息。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) AuthRegister(ctx context.Context, req *v1.AuthRegisterReq) (res *v1.AuthRegisterRes, err error) {
	blog.ControllerInfo(ctx, "AuthRegister", "用户注册")
	if !consts.SystemAbleRegisterValue {
		return nil, custom.ErrorSystemNotAbleRegister
	}
	iAuth := service.Auth()
	// 数据验证
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
	userInfo, errorCode := iAuth.UserRegister(ctx, req)
	if errorCode != nil {
		return nil, errorCode
	}

	return &v1.AuthRegisterRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "注册成功", userInfo),
	}, nil
}
