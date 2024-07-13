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
	"bamboo-service/internal/constant"
	"bamboo-service/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"

	"bamboo-service/api/auth/v1"
)

// AuthInitial
//
// # 用户初始化
//
// 用户初始化，用于用户初始化，需要传递用户名、手机号、密码、邮箱；
// 用作系统初始化的用户，为超级管理员账户；
// 拥有整个系统最高权限的账户；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - req			请求(*v1.AuthInitialReq)
//
// # 返回
//   - res			响应(*v1.AuthInitialRes)
//   - err			错误信息(error)
func (c *ControllerV1) AuthInitial(
	ctx context.Context,
	req *v1.AuthInitialReq,
) (res *v1.AuthInitialRes, err error) {
	g.Log().Notice(ctx, "[CONT] 用户初始化")
	// 是否为初始化模式
	if !constant.InitializeMode {
		return nil, berror.NewError(bcode.OperationNotAllow, "系统已经初始化，无法再次初始化系统")
	}
	// 获取 Header 中的 Referer
	err = service.Init().SetSystemWeb(ctx, req.Referer)
	if err != nil {
		return nil, err
	}
	// 初始化管理员
	err = service.Init().SetSystemAdmin(ctx, req.Username, req.Phone, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
