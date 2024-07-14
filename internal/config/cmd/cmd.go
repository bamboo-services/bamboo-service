/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package cmd

import (
	"bamboo-service/internal/config/middleware"
	"bamboo-service/internal/config/startup"
	"bamboo-service/internal/controller/auth"
	"bamboo-service/internal/controller/avatar"
	"bamboo-service/internal/controller/info"
	"bamboo-service/internal/controller/sms"
	"context"
	"github.com/bamboo-services/bamboo-utils/bmiddle"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "XiaoService 服务端",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化配置
			startup.SystemStartUp(ctx)

			// 服务器启动
			s := g.Server()
			// 关闭显示路由
			s.SetDumpRouterMap(false)

			// 路由组
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(bmiddle.BambooMiddleHandler)
				group.Middleware(middleware.MiddleSystemHasInitialized)

				// V1 版本路由
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Middleware(bmiddle.BambooMiddleDefaultCors)
					group.Middleware(middleware.MiddleHeaderCheck)
					group.Bind(
						auth.NewV1(),
						info.NewV1(),
					)
				})

				// V2 版本路由
				group.Group("/v2", func(group *ghttp.RouterGroup) {
					group.Middleware(bmiddle.BambooMiddleDefaultCors)
					group.Middleware(bmiddle.BambooMiddleRequestCheck)

					group.Bind(
						auth.NewV2(),
						sms.NewV2(),
						avatar.NewV2(),
					)
				})
			})

			// 特殊路由表
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					// 获取头像
					avatar.NewV3(),
				)
			})

			s.Run()
			return nil
		},
	}
)
