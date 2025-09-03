// @Author daixk 2024/8/2 23:26:00
package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
	"gtokenv2_test/internal/consts"
	"gtokenv2_test/internal/service"
	"gtokenv2_test/utility/rr"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

// CORS 跨域
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// HandlerResponseMiddleware 统一返回
func (s *sMiddleware) HandlerResponseMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err = r.GetError()
		res = r.GetHandlerResponse()
	)

	if err != nil {
		g.Log().Printf(r.Context(), "%+v", err)

		if _, ok := err.(gvalid.Error); ok {
			rr.FailedJsonWithCodeAndMessageExitAll(r, consts.CodeParamError, err.Error())
			return
		}

		rr.FailedJsonWithCodeAndMessageExitAll(r, consts.CodeServerError, consts.ErrCodeMessageMap[consts.CodeServerError])
		return
	}

	r.Response.WriteJsonExit(res)
}
