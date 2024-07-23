// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ip

import (
	"context"

	"bamboo-service/api/ip/v1"
)

type IIpV1 interface {
	IPImportIpv4(ctx context.Context, req *v1.IPImportIpv4Req) (res *v1.IPImportIpv4Res, err error)
	IPImportIpv6(ctx context.Context, req *v1.IPImportIpv6Req) (res *v1.IPImportIpv6Res, err error)
	IPUploadIPv4(ctx context.Context, req *v1.IPUploadIPv4Req) (res *v1.IPUploadIPv4Res, err error)
	IPUploadIPv6(ctx context.Context, req *v1.IPUploadIPv6Req) (res *v1.IPUploadIPv6Res, err error)
}
