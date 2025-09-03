// @Author daixk 2025/6/23 15:06:00
package websocket

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "gtokenv2_test/api/websocket/v1"
	"gtokenv2_test/gtoken"
	"gtokenv2_test/internal/consts"
	"gtokenv2_test/util/websockets"
	"gtokenv2_test/utility/rr"
)

func (c *cWebsocket) Websocket(ctx context.Context, req *v1.WebsocketReq) (res *rr.CommonRes, err error) {
	g.Log().Info(ctx, "调用websocket接口"+gtime.Now().String())

	token, err := gtoken.GetRequestToken(g.RequestFromCtx(ctx))
	if err != nil {
		return rr.FailedWithCodeAndMessage(consts.CodeNotAuthorized, consts.ErrCodeMessageMap[consts.CodeNotAuthorized]), nil
	}
	userKey, _, err := gtoken.GToken.ParseToken(ctx, token)
	if err != nil {
		return rr.FailedWithCodeAndMessage(consts.CodeNotAuthorized, consts.ErrCodeMessageMap[consts.CodeNotAuthorized]), nil
	}
	g.Dump(userKey)

	// 最大连接数
	if websockets.WsManage.GetWSCount() >= 1000000 {
		return rr.FailedWithCodeAndMessage(consts.CodeServerError, consts.ErrCodeMessageMap[consts.CodeServerError]), err
	}

	conn, err := upGrande.Upgrade(g.RequestFromCtx(ctx).Response.Writer, g.RequestFromCtx(ctx).Request, nil)
	if err != nil {
		return nil, err
	}

	websockets.WsManage.RegisterClient(ctx, userKey, conn)

	return rr.Success(), nil
}
