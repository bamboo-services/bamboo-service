package proxy

import (
	"bamboo-service/api/proxy/v1"
	"bamboo-service/internal/consts"
	"bamboo-service/internal/custom"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ProxyGroupAdd 添加新的代理组。
//
// 参数:
//   - ctx: 上下文对象，用于传递操作范围和数据。
//   - req: 包含代理组名称和描述的请求对象。
//
// 返回:
//   - res: 包含操作结果的响应对象。
//   - err: 可能的错误结果。
//
// 错误:
//   - 如果获取用户信息失败或用户不存在，将返回相应错误。
//   - 如果代理组创建失败，将返回相应错误。
func (c *ControllerV1) ProxyGroupAdd(ctx context.Context, req *v1.ProxyGroupAddReq) (res *v1.ProxyGroupAddRes, err error) {
	blog.ControllerInfo(ctx, "ProxyGroupAdd", "添加代理组")
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

	// 创建代理组
	iProxy := service.Proxy()
	_, errorCode = iProxy.CreateProxyGroup(ctx, getUser, req.Name, req.Description)
	if errorCode != nil {
		return nil, errorCode
	}

	return &v1.ProxyGroupAddRes{
		ResponseDTO: bresult.Success(ctx, "代理组创建成功"),
	}, nil
}
