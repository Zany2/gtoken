// @Author daixk 2024/8/29 0:09:00
package rr

// CommonReq 公共分页请求参数
type CommonReq struct {
	PageNum  int64 `json:"page_num" in:"query" d:"1"`   // 页码
	PageSize int64 `json:"page_size" in:"query" d:"30"` // 页记录数
}
