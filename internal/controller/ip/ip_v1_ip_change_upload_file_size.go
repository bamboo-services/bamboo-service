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

	"bamboo-service/api/ip/v1"
)

// IPChangeUploadFileSize
//
// # 修改上传文件大小
//
// 修改上传文件大小，用于修改上传文件大小操作；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - req			请求(*v1.IPChangeUploadFileSizeReq)
//
// # 返回
//   - res			响应(*v1.IPChangeUploadFileSizeRes)
//   - err			错误信息(error)
func (c *ControllerV1) IPChangeUploadFileSize(
	ctx context.Context,
	req *v1.IPChangeUploadFileSizeReq,
) (res *v1.IPChangeUploadFileSizeRes, err error) {
	g.Log().Notice(ctx, "[CONT] 修改上传文件大小")
	// 管理员授权认证
	err = service.Auth().CheckUserHasSuperAdmin(ctx, req.Authorization)
	if err != nil {
		return nil, err
	}
	// 修改上传文件大小
	g.Server().SetClientMaxBodySize(req.Length)
	return nil, nil
}
