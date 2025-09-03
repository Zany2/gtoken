package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"gtokenv2_test/gtoken"
	"gtokenv2_test/internal/controller/common"
	"gtokenv2_test/internal/controller/hello"
	"gtokenv2_test/internal/controller/user"
	"gtokenv2_test/internal/controller/websocket"
	"gtokenv2_test/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS, // 跨域
					gtoken.NewDefaultMiddleware(gtoken.GToken, gtoken.ResFunc).Auth, // gToken
					service.Middleware().HandlerResponseMiddleware,                  // 统一返回
				)

				// hello
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(
						hello.Hello,
					)
				})

				// websocket
				group.Group("/websocket", func(group *ghttp.RouterGroup) {
					group.Bind(
						websocket.Websocket,
					)
				})

				// common
				group.Group("/common", func(group *ghttp.RouterGroup) {
					group.Bind(
						common.Common,
					)
				})

				// user
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.User,
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
