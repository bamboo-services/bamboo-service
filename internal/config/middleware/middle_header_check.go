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
	"strings"
)

// MiddleHeaderCheck
//
// # 中间件-请求头检查
//
// 用于检查请求头中的 Referer 信息是否合法；
//
// # 参数
//   - r		请求(*ghttp.Request)
func MiddleHeaderCheck(r *ghttp.Request) {
	// 获取请求头信息
	referer := r.Header.Get("Referer")
	if referer == "" {
		r.SetError(berror.NewError(bcode.Unauthorized, "请求头中缺少 Referer 信息"))
		r.Response.Status = 401
		return
	}
	if !strings.Contains(constant.SystemReferer, referer) {
		r.SetError(berror.NewError(bcode.Forbidden, "Referer 信息不合法"))
		r.Response.Status = 403
		return
	}
	r.Middleware.Next()
}
