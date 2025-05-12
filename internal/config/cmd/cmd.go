package cmd

import (
	"bamboo-service/internal/controller/auth"
	"bamboo-service/internal/handler/hook"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/bhandler/bhook"
	"github.com/XiaoLFeng/bamboo-utils/bhandler/bmiddle"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 绑定事件
			s.BindHookHandler("/api/*", ghttp.HookBeforeServe, bhook.BambooHookDefaultCors)
			s.BindHookHandler("/api/*/admin/*", ghttp.HookBeforeServe, hook.BeforeAdminCheckHook)

			// API 接口
			s.Group("/api", func(api *ghttp.RouterGroup) {
				api.Middleware(bmiddle.BambooHandlerResponse)
				api.Group("/v1", func(v1 *ghttp.RouterGroup) {
					v1.Bind(
						auth.NewV1(),
					)
				})
			})

			// 静态资源和前端
			s.Group("/", func(route *ghttp.RouterGroup) {
				route.Bind()
			})

			s.Run()
			return nil
		},
	}
)
