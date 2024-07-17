/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package ip

import (
	"bamboo-service/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"bamboo-service/api/ip/v1"
)

// IPUploadIPv6
//
// # 上传IPv6数据库
//
// 上传IPv6数据库，用于上传IPv6数据库操作；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - req		请求(*v1.IPUploadIPv6Req)
//
// # 返回
//   - res		响应(*v1.IPUploadIPv6Res)
//   - err		错误信息(error)
func (c *ControllerV1) IPUploadIPv6(
	ctx context.Context,
	req *v1.IPUploadIPv6Req,
) (res *v1.IPUploadIPv6Res, err error) {
	g.Log().Notice(ctx, "[CONT] 上传IPv6数据库")
	// 授权检查
	err = service.Auth().CheckUserHasSuperAdmin(ctx, req.Authorization)
	if err != nil {
		return nil, err
	}
	getRequest := ghttp.RequestFromCtx(ctx)
	file := getRequest.GetUploadFile("file")
	// 上传文件
	err = service.IP().IPv6FileUpload(ctx, file)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
