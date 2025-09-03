// @Author daixk 2025/6/7 20:43:00
package common

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "gtokenv2_test/api/common/v1"
	"gtokenv2_test/utility/rr"
)

func (c *cCommon) SysConfig(ctx context.Context, req *v1.SysConfigReq) (res *rr.CommonRes, err error) {
	g.Log().Info(ctx, "test")
	return rr.SuccessWithMessage("这是测试系统"), err
}
