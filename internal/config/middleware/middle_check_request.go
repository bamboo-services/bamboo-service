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
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// MiddleRequestHandler
//
// # 中间件请求处理
//
// 中间件请求处理，用于处理请求的中间件；主要用于获取请求中的参数信息；用于测试以及调试；
//
// # 参数
//   - r		请求对象(*ghttp.Request)
func MiddleRequestHandler(r *ghttp.Request) {
	// 获取请求的参数
	g.Log().Noticef(r.Context(), "[REQU] 请求信息：[%s] %s", r.Method, r.URL.Path)
	if len(r.GetBody()) > 0 {
		g.Log().Infof(r.Context(), "\t[BODY]请求体参数:")
		// 解析请求体参数
		decode, _ := gjson.Decode(r.GetBody())
		for key, value := range decode.(map[string]interface{}) {
			g.Log().Infof(r.Context(), "\t\t[%v] \t%v", key, value)
		}
	}
	if len(r.GetQueryMap()) > 0 {
		g.Log().Infof(r.Context(), "\t[PARA]请求参数:")
		for key, value := range r.GetQueryMap() {
			g.Log().Infof(r.Context(), "\t\t[%v] \t%v", key, value)
		}
	}
	if r.Request.Header != nil {
		g.Log().Debugf(r.Context(), "\t[HEAD]请求头部:")
		for key, value := range r.Request.Header {
			g.Log().Debugf(r.Context(), "\t\t[%v] \t%v", key, value)
		}
	}

	startTime := gtime.Now().UnixMilli()
	r.Middleware.Next()
	endTime := gtime.Now().UnixMilli()

	// 请求时间统计
	g.Log().Noticef(r.Context(), "[TIME] 请求耗时：%d ms", endTime-startTime)
}
