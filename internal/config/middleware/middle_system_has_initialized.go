/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package middleware

import (
	"bamboo-service/internal/constant"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// MiddleSystemHasInitialized
//
// # 系统初始化
//
// 用于检查系统是否进行初始化；
// 若没有进行初始化则将系统进行拦截，只放行初始化接口；
//
// # 参数
//   - r		*ghttp.Request		请求对象
func MiddleSystemHasInitialized(r *ghttp.Request) {
	// 检查是否已初始化
	if constant.InitializeMode {
		// 检查访问接口是否是初始化接口
		if r.Request.URL.Path == "/api/v1/auth/init" {
			r.Middleware.Next()
		} else {
			r.SetError(berror.NewError(bcode.OperationNotAllow, "系统未初始化，不允许访问其他接口"))
		}
	} else {
		r.Middleware.Next()
	}
}
