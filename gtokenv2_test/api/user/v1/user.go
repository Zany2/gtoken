// Author daixk 2023-12-01 15:06:47
package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	UserName string `json:"user_name" v:"required#用户名不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type UserInfoReq struct {
	g.Meta `path:"/info" method:"get"`
}
