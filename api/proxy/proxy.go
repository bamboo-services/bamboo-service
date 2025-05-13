// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package proxy

import (
	"context"

	"bamboo-service/api/proxy/v1"
)

type IProxyV1 interface {
	ProxyTokenGenerate(ctx context.Context, req *v1.ProxyTokenGenerateReq) (res *v1.ProxyTokenGenerateRes, err error)
}
