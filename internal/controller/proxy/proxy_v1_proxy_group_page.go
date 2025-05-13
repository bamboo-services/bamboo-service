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

// ProxyGroupPage 获取代理组列表（分页）。
//
// 参数:
//   - ctx: 上下文对象，用于传递操作范围和数据。
//   - req: 包含分页信息和搜索内容的请求对象。
//
// 返回:
//   - res: 包含分页后代理组数据的响应对象。
//   - err: 可能的错误结果。
//
// 错误:
//   - 若用户信息获取失败或用户不存在，将返回对应错误。
//   - 若代理组列表查询失败，将返回对应错误。
func (c *ControllerV1) ProxyGroupPage(ctx context.Context, req *v1.ProxyGroupPageReq) (res *v1.ProxyGroupPageRes, err error) {
	blog.ControllerInfo(ctx, "ProxyGroupPage", "代理组列表（分页）")
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

	// 获取代理组列表
	iProxy := service.Proxy()
	proxyGroupList, errorCode := iProxy.ProxyGroupPage(ctx, getUser, req.Page, req.Size, req.Search)
	if errorCode != nil {
		return nil, errorCode
	}
	return &v1.ProxyGroupPageRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取代理组列表成功", proxyGroupList),
	}, nil
}
