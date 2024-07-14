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
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"bamboo-service/api/avatar/v3"
)

func (c *ControllerV3) AvatarGetImage(ctx context.Context, req *v3.AvatarGetImageReq) (res *v3.AvatarGetImageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
