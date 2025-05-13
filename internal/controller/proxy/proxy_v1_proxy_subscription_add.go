package proxy

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/custom"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/net/ghttp"

	"bamboo-service/api/proxy/v1"
)

// ProxySubscriptionAdd 向指定代理组添加订阅内容。
//
// 参数:
//   - ctx: 上下文对象，用于传递操作范围和数据。
//   - req: 包含代理组UUID、订阅名称、商户名称、订阅描述和URL的请求对象。
//
// 返回:
//   - res: 包含操作结果的响应对象。
//   - err: 可能的错误结果。
//
// 错误:
//   - 用户信息获取失败或用户不存在。
//   - 订阅添加过程中发生错误。
func (c *ControllerV1) ProxySubscriptionAdd(ctx context.Context, req *v1.ProxySubscriptionAddReq) (res *v1.ProxySubscriptionAddRes, err error) {
	blog.ControllerInfo(ctx, "ProxySubscriptionAdd", "添加订阅")
	// 获取用户信息
	request := ghttp.RequestFromCtx(ctx)
	iUser := service.User()
	getUser, errorCode := iUser.GetUserByUUID(ctx, request.GetHeader(consts.HeaderUserUUID))
	if errorCode != nil {
		return nil, errorCode
	}
	if getUser == nil {
		return nil, custom.ErrorUserNotExist
	}

	// 添加订阅
	iProxy := service.Proxy()
	_, errorCode = iProxy.AddSubscriptionInProxyGroup(ctx, getUser, req.ProxyGroupUuid, req.Name, req.Merchant, req.Description, req.Url)
	if errorCode != nil {
		return nil, errorCode
	}
	return &v1.ProxySubscriptionAddRes{
		ResponseDTO: bresult.Success(ctx, "订阅添加成功"),
	}, nil
}
