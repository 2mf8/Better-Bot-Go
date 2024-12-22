package websocket

import (
	"github.com/2mf8/Better-Bot-Go/util"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// safe websocket
type SafeWebSocket struct {
	Conn          *websocket.Conn
	SendChannel   chan *WebSocketSendingMessage
	OnRecvMessage func(messageType int, data []byte)
	OnClose       func(int, string)
}

type WebSocketSendingMessage struct {
	MessageType int
	Data        []byte
}

func (ws *SafeWebSocket) Send(messageType int, data []byte) {
	ws.SendChannel <- &WebSocketSendingMessage{
		MessageType: messageType,
		Data:        data,
	}
}

func NewSafeWebSocket(conn *websocket.Conn, OnRecvMessage func(messageType int, data []byte), onClose func(int, string)) *SafeWebSocket {
	ws := &SafeWebSocket{
		Conn:          conn,
		SendChannel:   make(chan *WebSocketSendingMessage, 100),
		OnRecvMessage: OnRecvMessage,
		OnClose:       onClose,
	}

	conn.SetCloseHandler(func(code int, text string) error {
		ws.OnClose(code, text)
		return nil
	})

	// 接受消息
	util.SafeGo(func() {
		for {
			messageType, data, err := conn.ReadMessage()
			if err != nil {
				log.Errorf("failed to read message, err: %+v", err)
				_ = conn.Close()
				return
			}
			if messageType == websocket.PingMessage {
				ws.Send(websocket.PongMessage, []byte("pong"))
				continue
			}
			ws.OnRecvMessage(messageType, data)
		}
	})

	// 发送消息
	util.SafeGo(func() {
		for sendingMessage := range ws.SendChannel {
			if ws.Conn == nil {
				log.Errorf("failed to send websocket message, conn is nil")
				return
			}
			err := ws.Conn.WriteMessage(sendingMessage.MessageType, sendingMessage.Data)
			if err != nil {
				log.Errorf("failed to send websocket message, %+v", err)
				_ = conn.Close()
				return
			}
		}
	})
	return ws
}
