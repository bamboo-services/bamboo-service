// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package proxy

import (
	"context"

	"bamboo-service/api/proxy/v1"
)

type IProxyV1 interface {
	ProxyGroupAdd(ctx context.Context, req *v1.ProxyGroupAddReq) (res *v1.ProxyGroupAddRes, err error)
	ProxyGroupPage(ctx context.Context, req *v1.ProxyGroupPageReq) (res *v1.ProxyGroupPageRes, err error)
	ProxySubscriptionAdd(ctx context.Context, req *v1.ProxySubscriptionAddReq) (res *v1.ProxySubscriptionAddRes, err error)
	ProxyTokenGenerate(ctx context.Context, req *v1.ProxyTokenGenerateReq) (res *v1.ProxyTokenGenerateRes, err error)
}
