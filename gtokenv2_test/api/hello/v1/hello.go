// @Author daixk 2025/6/7 21:42:00
package v1

import "github.com/gogf/gf/v2/frame/g"

type HelloReq struct {
	g.Meta `path:"/hello" method:"get"`
}
