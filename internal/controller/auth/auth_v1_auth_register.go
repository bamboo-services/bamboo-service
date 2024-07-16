/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"bamboo-service/api/auth/v1"
)

// AuthRegister
//
// # 用户注册
//
// 用户注册，用于用户注册操作；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - req			请求(*v1.AuthRegisterReq)
//
// # 返回
//   - res			响应(*v1.AuthRegisterRes)
//   - err			错误信息(error)
func (c *ControllerV1) AuthRegister(
	ctx context.Context,
	req *v1.AuthRegisterReq,
) (res *v1.AuthRegisterRes, err error) {
	g.Log().Notice(ctx, "[CONT] 用户注册")

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
