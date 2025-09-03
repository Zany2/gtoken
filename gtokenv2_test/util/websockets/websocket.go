package websockets

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

var (
	once      sync.Once
	WsManage  *WebSocketManager
	WsHandler *WsHandlerFunc
)

func init() {
	once.Do(func() {
		WsHandler = &WsHandlerFunc{
			UserCtxMap:    make(map[string]context.Context),
			UserCtxCancel: make(map[string]context.CancelFunc),
		}
		WsManage = InitWebSocketManager(WsHandler)
	})
}

type WsMessageRes struct {
	Data string `json:"data"`
}

type WsHandlerFunc struct {
	UserCtxMap    map[string]context.Context
	UserCtxCancel map[string]context.CancelFunc
}

// SendMessage 发送消息
func SendMessage(uniqueCode string, in *WsMessageRes) {
	if in == nil {
		return
	}
	jb, _ := json.Marshal(in)
	WsManage.SendMessage(uniqueCode, jb)
}

// OnMessage 收到信息
func (ws *WsHandlerFunc) OnMessage(UniqueCode string, message []byte) {

	fmt.Println(UniqueCode)
	fmt.Println(string(message))

}

// OnOpen 建立连接
func (ws *WsHandlerFunc) OnOpen(UniqueCode string) {
	//WsManage.SendMessage(UniqueCode, []byte(UniqueCode))

	ctx := WsManage.GetClientCtx(UniqueCode)
	ctx, cancel := context.WithCancel(ctx)
	ws.UserCtxCancel[UniqueCode] = cancel

}

// OnClose 断开连接
func (ws *WsHandlerFunc) OnClose(UniqueCode string) {
	ws.UserCtxCancel[UniqueCode]()
}
