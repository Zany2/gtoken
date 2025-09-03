// Author daixk 2023-12-04 09:06:38
package rr

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"gtokenv2_test/internal/consts"
)

// CommonRes 数据返回通用JSON数据结构
type CommonRes struct {
	Code    int         `json:"code"`    // 提示码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// CommonResModel 数据数据返回体
type CommonResModel[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"`
}

// Json 返回标准JSON数据
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	r.Response.WriteJson(CommonRes{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// SuccessJsonExitAll 成功返回
func SuccessJsonExitAll(r *ghttp.Request) {
	Json(r, consts.CodeOK, consts.ErrCodeMessageMap[consts.CodeOK], g.Array{})
	r.ExitAll()
}

// SuccessJsonWithMessageExitAll 带有提示的成功返回
func SuccessJsonWithMessageExitAll(r *ghttp.Request, message string) {
	Json(r, consts.CodeOK, message, g.Array{})
	r.ExitAll()
}

// SuccessJsonWithDataExitAll 带有数据的成功返回
func SuccessJsonWithDataExitAll(r *ghttp.Request, data interface{}) {
	Json(r, consts.CodeOK, consts.ErrCodeMessageMap[consts.CodeOK], data)
	r.ExitAll()
}

// SuccessJsonWithMessageAndData 带有提示和数据的成功返回
func SuccessJsonWithMessageAndData(r *ghttp.Request, message string, data interface{}) {
	Json(r, consts.CodeOK, message, data)
	r.ExitAll()
}

// FailedJsonExitAll 失败返回
func FailedJsonExitAll(r *ghttp.Request) {
	Json(r, consts.CodeServerError, consts.ErrCodeMessageMap[consts.CodeServerError], g.Array{})
	r.ExitAll()
}

// FailedJsonWithMessageExitAll 带有提示的失败返回
func FailedJsonWithMessageExitAll(r *ghttp.Request, message string) {
	Json(r, consts.CodeServerError, message, g.Array{})
	r.ExitAll()
}

// FailedJsonWithCodeAndMessageExitAll 带有提示码和提示的失败返回
func FailedJsonWithCodeAndMessageExitAll(r *ghttp.Request, code int, message string) {
	Json(r, code, message, g.Array{})
	r.ExitAll()
}

// Success 统一成功返回
func Success() *CommonRes {
	return &CommonRes{
		Code:    consts.CodeOK,
		Message: consts.ErrCodeMessageMap[consts.CodeOK],
		Data:    g.Array{},
	}
}

// SuccessWithMessage 带有提示的统一成功返回
func SuccessWithMessage(message string) *CommonRes {
	return &CommonRes{
		Code:    consts.CodeOK,
		Message: message,
		Data:    g.Array{},
	}
}

// SuccessWithData 带有数据的统一成功返回
func SuccessWithData(data interface{}) *CommonRes {
	return &CommonRes{
		Code:    consts.CodeOK,
		Message: consts.ErrCodeMessageMap[consts.CodeOK],
		Data:    data,
	}
}

// SuccessWithMessageAndData 带有提示和数据的统一成功返回
func SuccessWithMessageAndData(message string, data interface{}) *CommonRes {
	return &CommonRes{
		Code:    consts.CodeOK,
		Message: message,
		Data:    data,
	}
}

// Failed 统一失败返回
func Failed() *CommonRes {
	return &CommonRes{
		Code:    consts.CodeServerError,
		Message: consts.ErrCodeMessageMap[consts.CodeServerError],
		Data:    g.Array{},
	}
}

// FailedWithMessage 带有提示的统一失败返回
func FailedWithMessage(message string) *CommonRes {
	return &CommonRes{
		Code:    consts.CodeServerError,
		Message: message,
		Data:    g.Array{},
	}
}

// FailedWithCodeAndMessage 带有提示码和提示的统一失败返回
func FailedWithCodeAndMessage(code int, message string) *CommonRes {
	return &CommonRes{
		Code:    code,
		Message: message,
		Data:    g.Array{},
	}
}

// FailedWithMessageAndData 带有提示和数据的统一失败返回
func FailedWithMessageAndData(message string, data interface{}) *CommonRes {
	return &CommonRes{
		Code:    consts.CodeServerError,
		Message: message,
		Data:    data,
	}
}
