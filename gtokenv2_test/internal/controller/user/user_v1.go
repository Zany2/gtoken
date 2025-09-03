// @Author daixk 2025/6/7 20:31:00
package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gtokenv2_test/api/user/v1"
	"gtokenv2_test/gtoken"
	"gtokenv2_test/internal/consts"
	"gtokenv2_test/util/websockets"
	"gtokenv2_test/utility/rr"
)

func (c *cUser) Login(ctx context.Context, req *v1.LoginReq) (res *rr.CommonRes, err error) {
	// TODO 查询数据库操作 拼接用户信息 基本信息、权限信息......

	if req.UserName != "daixk" {
		return rr.FailedWithCodeAndMessage(consts.CodeServerError, "该用户不存在"), err
	}
	if req.UserName == "daixk" && req.Password != "daixk" {
		return rr.FailedWithCodeAndMessage(consts.CodeServerError, "密码错误"), err
	}

	token, err := gtoken.GToken.Generate(
		ctx,
		req.UserName,
		g.Map{
			consts.CtxUserId:           req.UserName,
			consts.CtxOrganizationUuid: req.UserName,
		})
	if err != nil {
		return nil, err
	}

	return rr.SuccessWithData(g.Map{
		"token": token,
	}), err
}

func (c *cUser) Logout(ctx context.Context, req *v1.LogoutReq) (res *rr.CommonRes, err error) {
	token, err := gtoken.GetRequestToken(g.RequestFromCtx(ctx))
	if err != nil {
		return
	}
	userKey, _, err := gtoken.GToken.ParseToken(ctx, token)
	if err != nil {
		return
	}
	err = gtoken.GToken.Destroy(ctx, userKey)
	return rr.Success(), err
}

func (c *cUser) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *rr.CommonRes, err error) {
	token, err := gtoken.GetRequestToken(g.RequestFromCtx(ctx))
	if err != nil {
		return
	}
	userKey, data, err := gtoken.GToken.ParseToken(ctx, token)
	if err != nil {
		return
	}
	ctxValueMap := gconv.Map(ctx.Value(gtoken.KeyUserKey))

	g.Dump("------------------------------------------------")
	g.Dump("这是直接获取的上下文ctx中以userKey存储的value")
	g.Dump(ctx.Value(gtoken.KeyUserKey))
	g.Dump("------------------------------------------------")
	g.Dump("这是以上下文ctx和token获取的缓存的data")
	g.Dump(data)
	g.Dump("------------------------------------------------")
	g.Dump(ctxValueMap[consts.CtxUserId])
	g.Dump(ctxValueMap[consts.CtxUserId])
	g.Dump("------------------------------------------------")

	websockets.SendMessage(userKey, &websockets.WsMessageRes{Data: "111111111111111111111111111111111111111111111111"})

	return rr.Success(), err
}
