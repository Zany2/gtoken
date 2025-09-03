package websockets

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"sync"
	"sync/atomic"
	"time"
)

type HandlerFunc interface {
	OnOpen(UniqueCode string)
	OnMessage(UniqueCode string, message []byte)
	OnClose(UniqueCode string)
}

type WebSocketManager struct {
	count       int64
	clientGroup sync.Map // map[string]*Client
	Handler     HandlerFunc
}

// InitWebSocketManager 初始化WebSocketManager
func InitWebSocketManager(handler HandlerFunc) *WebSocketManager {
	wsManager := &WebSocketManager{Handler: handler}
	go wsManager.heartbeat()
	return wsManager
}

// heartbeat 保活心跳
func (m *WebSocketManager) heartbeat() {
	// 创建一个新的 ticker，每 15 秒触发一次
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop() // 停止 ticker 防止泄露

	for {
		select {
		case <-ticker.C:
			// 模拟每个客户端发送 Ping 消息
			m.clientGroup.Range(func(key, value interface{}) bool {
				value.(*Client).message <- &messageData{data: nil, dataType: websocket.PingMessage}
				return true
			})
		}
	}
}

// onConnect 新建连接
func (m *WebSocketManager) onConnect(client *Client) {
	if g.IsEmpty(client.uniqueCode) {
		client.message <- &messageData{data: nil, dataType: websocket.CloseMessage}
		return
	}

	atomic.AddInt64(&m.count, 1)
	m.safeAddClientClientGroup(client)
	m.Handler.OnOpen(client.uniqueCode)
}

// onMessage 收到信息
func (m *WebSocketManager) onMessage(client *Client, message []byte) {
	m.Handler.OnMessage(client.uniqueCode, message)
}

// disConnect 关闭conn
func (m *WebSocketManager) disConnect(client *Client) {
	atomic.AddInt64(&m.count, -1)
	m.safeRemoveClientGroup(client.uniqueCode)
	m.Handler.OnClose(client.uniqueCode)
}

// safeRemoveClientGroup 移除conn
func (m *WebSocketManager) safeRemoveClientGroup(UniqueCode string) {
	if client, ok := m.clientGroup.LoadAndDelete(UniqueCode); ok {
		close(client.(*Client).message)
	}
}

// safeAddClientClientGroup 新增conn
func (m *WebSocketManager) safeAddClientClientGroup(client *Client) {
	m.clientGroup.Store(client.uniqueCode, client)
}

// RegisterClient 注册
func (m *WebSocketManager) RegisterClient(ctx context.Context, UniqueCode string, conn *websocket.Conn) {
	client := &Client{Ctx: ctx, uniqueCode: UniqueCode, conn: conn, manager: m, message: make(chan *messageData)}
	go client.Read()
	go client.Write()
	m.onConnect(client)
}

// GetClientCtx 获取ctx
func (m *WebSocketManager) GetClientCtx(UniqueCode string) context.Context {
	if client, ok := m.clientGroup.Load(UniqueCode); ok {
		return client.(*Client).Ctx
	}
	return nil
}

// GetWSCount 获取连接数
func (m *WebSocketManager) GetWSCount() int64 {
	return atomic.LoadInt64(&m.count)
}

// SendMessage 为指定客户端发送消息 UTF-8 编码
// UniqueCode 连接的唯一标识
// message 需要发送的消息
func (m *WebSocketManager) SendMessage(UniqueCode string, message []byte) {
	if client, ok := m.clientGroup.Load(UniqueCode); ok {
		client.(*Client).message <- &messageData{data: message, dataType: websocket.TextMessage}
	}
}

// BroadcastMessage 广播消息 UTF-8 编码
// message 需要广播的消息
// excluded 排除广播的连接的唯一标识
func (m *WebSocketManager) BroadcastMessage(message []byte, excluded ...string) {
	excludedMap := make(map[string]struct{}, len(excluded))
	for i := range excluded {
		excludedMap[excluded[i]] = struct{}{}
	}

	m.clientGroup.Range(func(key, value interface{}) bool {
		if _, ok := excludedMap[key.(string)]; !ok {
			value.(*Client).message <- &messageData{data: message, dataType: websocket.TextMessage}
		}
		return true
	})
}

// CloseWebsocketConn 关闭指定客户端得连接
// UniqueCode 连接的唯一标识
func (m *WebSocketManager) CloseWebsocketConn(UniqueCode string) {
	if client, ok := m.clientGroup.Load(UniqueCode); ok {
		client.(*Client).message <- &messageData{data: nil, dataType: websocket.CloseMessage}
	}
}

// SendBinary 为指定客户端发送二进制数据
// UniqueCode 连接的唯一标识
// binary 需要发送的二进制数据
func (m *WebSocketManager) SendBinary(UniqueCode string, binary []byte) {
	if client, ok := m.clientGroup.Load(UniqueCode); ok {
		client.(*Client).message <- &messageData{data: binary, dataType: websocket.BinaryMessage}
	}
}

// BroadcastBinary 广播二进制数据
// message 需要广播的二进制数据
// excluded 排除广播的连接的唯一标识
func (m *WebSocketManager) BroadcastBinary(binary []byte, excluded ...string) {
	excludedMap := make(map[string]struct{}, len(excluded))
	for i := range excluded {
		excludedMap[excluded[i]] = struct{}{}
	}

	m.clientGroup.Range(func(key, value interface{}) bool {
		if _, ok := excludedMap[key.(string)]; !ok {
			value.(*Client).message <- &messageData{data: binary, dataType: websocket.BinaryMessage}
		}
		return true
	})
}
