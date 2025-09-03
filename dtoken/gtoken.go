// @Author daixk 2025/6/7 22:50:00
package gtoken

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"gtokenv2_test/internal/consts"
	"gtokenv2_test/utility/rr"
)

var (
	GToken  Token
	ResFunc = func(r *ghttp.Request) {
		rr.FailedJsonWithCodeAndMessageExitAll(r, consts.CodeNotAuthorized, consts.ErrCodeMessageMap[consts.CodeNotAuthorized])
	}
)

func init() {
	GToken = NewDefaultTokenByConfig()
}
