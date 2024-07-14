/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package avatar

import (
	"bamboo-service/internal/service"
	"bamboo-service/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"bamboo-service/api/avatar/v3"
)

// AvatarGetImageNoHash
//
// # 获取头像
//
// 获取头像，获取用户头像；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - req		请求参数(*v3.AvatarGetImageNoHashReq)
//
// # 返回
//   - res		返回结果(*v3.AvatarGetImageNoHashRes)
//   - err		错误信息(error)
func (c *ControllerV3) AvatarGetImageNoHash(
	ctx context.Context,
	req *v3.AvatarGetImageNoHashReq,
) (res *v3.AvatarGetImageNoHashRes, err error) {
	g.Log().Notice(ctx, "[CONT] 获取头像")
	getRequest := g.RequestFromCtx(ctx)
	avatar := service.Avatar().AvatarGet(ctx, utility.StringToMD5(req.Email))
	getRequest.Response.Write(avatar)
	return nil, nil
}
