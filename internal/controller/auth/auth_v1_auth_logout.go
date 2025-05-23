package auth

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/net/ghttp"

	"bamboo-service/api/auth/v1"
)

// AuthLogout 用户登出并注销授权令牌。
//
// 参数:
//   - ctx: 上下文信息。
//   - req: 用户登出请求，包含必要的请求头信息（用户 UUID 和 Token）。
//
// 返回:
//   - res: 用户登出响应，表示登出结果。
//   - err: 执行过程中可能发生的错误。
func (c *ControllerV1) AuthLogout(ctx context.Context, req *v1.AuthLogoutReq) (res *v1.AuthLogoutRes, err error) {
	blog.ControllerInfo(ctx, "AuthLogout", "用户登出")
	// 获取用户 UUID 和 Token
	request := ghttp.RequestFromCtx(ctx)
	getUserUUID := request.GetHeader(consts.HeaderUserUUID)
	getToken := request.GetHeader(consts.HeaderToken)

	if getUserUUID == "" || getToken == "" {
		return nil, berror.ErrorAddData(&berror.ErrInvalidParameters, "用户 UUID 或 Token 不能为空")
	}

	// 删除用户 Token「进行登出操作」
	iToken := service.Token()
	errorCode := iToken.RemoveToken(ctx, getUserUUID, getToken)
	if errorCode != nil {
		return nil, errorCode
	}
	return &v1.AuthLogoutRes{
		ResponseDTO: bresult.Success(ctx, "登出成功"),
	}, nil
}
