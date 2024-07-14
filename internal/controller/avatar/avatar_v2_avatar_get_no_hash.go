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

	"bamboo-service/api/avatar/v2"
)

func (c *ControllerV2) AvatarGetNoHash(
	ctx context.Context,
	req *v2.AvatarGetNoHashReq,
) (res *v2.AvatarGetNoHashRes, err error) {
	g.Log().Notice(ctx, "[CONT] 获取头像接口")
	avatar, err := service.Avatar().AvatarGetAPI(ctx, utility.StringToMD5(req.Email))
	if err != nil {
		return nil, err
	}
	return &v2.AvatarGetNoHashRes{
		AvatarApiDTO: *avatar,
	}, nil
}
