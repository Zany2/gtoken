// @Author daixk 2025/6/7 21:43:00
package hello

import (
	"context"
	v1 "gtokenv2_test/api/hello/v1"
	"gtokenv2_test/utility/rr"
)

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *rr.CommonRes, err error) {
	return rr.SuccessWithData("hello"), err
}
