// @Author daixk 2024/7/24 14:07:00
package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var Websocket = cWebsocket{}

type cWebsocket struct{}

var upGrande websocket.Upgrader

func init() {
	upGrande = websocket.Upgrader{
		// HandshakeTimeout specifies the duration for the handshake to complete.
		HandshakeTimeout: time.Second * 5,
		//设置允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
