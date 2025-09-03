package consts

const (
	CtxUserId           = "CtxUserId"           // 全局上下文CtxUserId
	CtxOrganizationUuid = "CtxOrganizationUuid" // 全局上下文CtxOrganizationUuid
)

// 全局Code状态码
const (
	CodeOK             = 20000 // 请求成功
	CodeFirstTimeLogin = 20001 // 首次登录，请修改密码（非必须）

	CodeBadRequest    = 40000 // 错误请求
	CodeNotAuthorized = 40001 // 认证信息token不存在或已过期
	CodeParamError    = 40002 // 参数错误
	CodeNotFound      = 40003 // 资源不存在

	CodeServerError     = 50000 // 服务异常
	CodeServerBusy      = 50001 // 服务繁忙
	CodeAccountDisabled = 50002 // 账户被禁用
	CodeNotPermission   = 50003 // 暂无权限
	CodeIPLimited       = 50004 // 访问IP已被禁用
	CodeRateLimit       = 50005 // 接口限流
)

var ErrCodeMessageMap map[int]string

func init() {
	ErrCodeMessageMap = make(map[int]string)

	// 全局CodeMessage
	ErrCodeMessageMap[CodeOK] = "请求成功"
	ErrCodeMessageMap[CodeFirstTimeLogin] = "首次登录请先修改密码"

	ErrCodeMessageMap[CodeBadRequest] = "请求错误"
	ErrCodeMessageMap[CodeNotAuthorized] = "认证过期，请重新登录"
	ErrCodeMessageMap[CodeParamError] = "请求参数有误"
	ErrCodeMessageMap[CodeNotFound] = "访问资源不存在"

	ErrCodeMessageMap[CodeServerError] = "服务异常，请稍后再试或联系管理员"
	ErrCodeMessageMap[CodeServerBusy] = "服务繁忙，请稍后再试或联系管理员"
	ErrCodeMessageMap[CodeAccountDisabled] = "该账户已被禁用"
	ErrCodeMessageMap[CodeNotPermission] = "暂无权限"
	ErrCodeMessageMap[CodeIPLimited] = "该IP暂时无法访问"
	ErrCodeMessageMap[CodeRateLimit] = "操作过快，请稍后再试"
}
