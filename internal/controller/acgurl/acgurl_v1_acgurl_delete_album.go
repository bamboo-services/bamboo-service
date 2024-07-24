/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package acgurl

import (
	"bamboo-service/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"

	"bamboo-service/api/acgurl/v1"
)

// AcgurlDeleteAlbum
//
// # 删除图库
//
// 删除一个图库，用于删除一个图库操作；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - req		请求(*v1.AcgurlDeleteAlbumReq)
//
// # 返回
//   - res		响应(*v1.AcgurlDeleteAlbumRes)
//   - err		错误(error)
func (c *ControllerV1) AcgurlDeleteAlbum(
	ctx context.Context,
	req *v1.AcgurlDeleteAlbumReq,
) (res *v1.AcgurlDeleteAlbumRes, err error) {
	g.Log().Notice(ctx, "[CONT] 删除图库")
	// 权限校验
	if err = service.Auth().CheckUserHasLogin(ctx, req.Authorization); err != nil {
		return nil, err
	}
	// 删除图库
	err = service.Acgurl().DeleteAlbum(ctx, butil.StringToUUID(req.AlbumUUID), req.Authorization)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
