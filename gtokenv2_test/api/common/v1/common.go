// @Author daixk 2025/6/7 20:33:00
package v1

import "github.com/gogf/gf/v2/frame/g"

type SysConfigReq struct {
	g.Meta `path:"/config" method:"get"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}
