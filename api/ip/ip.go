/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ip

import (
	"context"

	"bamboo-service/api/ip/v1"
)

type IIpV1 interface {
	IPChangeUploadFileSize(ctx context.Context, req *v1.IPChangeUploadFileSizeReq) (res *v1.IPChangeUploadFileSizeRes, err error)
	IPImportIpv4(ctx context.Context, req *v1.IPImportIpv4Req) (res *v1.IPImportIpv4Res, err error)
	IPImportIpv6(ctx context.Context, req *v1.IPImportIpv6Req) (res *v1.IPImportIpv6Res, err error)
	IPUploadIPv4(ctx context.Context, req *v1.IPUploadIPv4Req) (res *v1.IPUploadIPv4Res, err error)
	IPUploadIPv6(ctx context.Context, req *v1.IPUploadIPv6Req) (res *v1.IPUploadIPv6Res, err error)
}
