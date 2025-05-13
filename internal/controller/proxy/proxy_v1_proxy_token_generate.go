package proxy

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"bamboo-service/api/proxy/v1"
)

// ProxyTokenGenerate 生成用于代理授权的令牌。
//
// 参数:
//   - ctx: 上下文对象，用于传递操作范围和数据。
//   - req: 包含令牌名称、描述、过期时间等信息的请求对象。
//
// 返回:
//   - res: 包含生成令牌信息的响应对象。
//   - err: 可能的错误结果。
//
// 错误:
//   - 如果令牌有效期超过1年或者其他参数校验失败，将返回相应错误。
func (c *ControllerV1) ProxyTokenGenerate(ctx context.Context, req *v1.ProxyTokenGenerateReq) (res *v1.ProxyTokenGenerateRes, err error) {
	blog.ControllerInfo(ctx, "ProxyTokenGenerate", "生成代理授权令牌")

	// 获取用户 UUID 信息
	request := ghttp.RequestFromCtx(ctx)
	getUserUUID := request.GetHeader(consts.HeaderUserUUID)

	// 如果设置有效期时间不可以超过 1 年
	if req.ExpiredAt != nil && req.ExpiredAt.After(gtime.Now().AddDate(1, 0, 0)) {
		blog.ControllerError(ctx, "ProxyTokenGenerate", "代理令牌有效期不能超过1年")
		return nil, berror.ErrorAddData(&berror.ErrInvalidParameters, "代理令牌有效期不能超过1年")
	}

	// 生成令牌
	iToken := service.Token()
	newToken, errorCode := iToken.GenerateProxyToken(ctx, getUserUUID, req.Name, req.Description, req.ExpiredAt)
	if errorCode != nil {
		return nil, errorCode
	}

	return &v1.ProxyTokenGenerateRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "令牌生成成功", newToken),
	}, nil
}
