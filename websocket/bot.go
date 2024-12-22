package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/2mf8/Better-Bot-Go/onebot"
	"github.com/2mf8/Better-Bot-Go/util"
	"github.com/fanliao/go-promise"
	"github.com/gorilla/websocket"
)

var Bots = make(map[string]*Bot)
var echo = ""

type Bot struct {
	RemoteAddr    string
	Session       *SafeWebSocket
	mux           sync.RWMutex
	WaitingFrames map[string]*promise.Promise
}

func NewBot(conn *websocket.Conn) *Bot {
	addr := conn.RemoteAddr()
	messageHandler := func(messageType int, data []byte) {
		_, ok := Bots[addr.String()]
		if !ok {
			_ = conn.Close()
			return
		}
	}
	closeHandler := func(code int, message string) {
		fmt.Printf("WebSocket客户端 %s 已断开连接\n", addr)
		delete(Bots, addr.String())
	}
	safeWs := NewSafeWebSocket(conn, messageHandler, closeHandler)
	bot := &Bot{
		RemoteAddr:    addr.String(),
		Session:       safeWs,
		WaitingFrames: make(map[string]*promise.Promise),
	}
	Bots[addr.String()] = bot
	fmt.Printf("新的WebSocket客户端已连接：%s\n", bot.RemoteAddr)
	fmt.Println("所有WebSocket客户端地址列表：")
	for naddr, _ := range Bots {
		println(naddr)
	}
	return bot
}

func sendFrameAndWait(bot *Bot, appid string, frame *onebot.Frame) (*onebot.Frame, error) {
	frame.BotId = appid
	frame.Echo = util.GenerateIdStr()
	frame.Ok = true
	data, err := json.Marshal(frame)
	if err != nil {
		return nil, err
	}
	bot.Session.Send(websocket.BinaryMessage, data)
	p := promise.NewPromise()
	bot.setWaitingFrame(frame.Echo, p)
	defer bot.delWaitingFrame(frame.Echo)
	resp, err, timeout := p.GetOrTimeout(120000)
	if err != nil || timeout {
		return nil, err
	}
	respFrame, ok := resp.(*onebot.Frame)
	if !ok {
		return nil, errors.New("failed to convert promise result to resp frame")
	}
	return respFrame, nil
}

func (bot *Bot) setWaitingFrame(key string, value *promise.Promise) {
	bot.mux.Lock()
	defer bot.mux.Unlock()
	bot.WaitingFrames[key] = value
}

func (bot *Bot) getWaitingFrame(key string) (*promise.Promise, bool) {
	bot.mux.RLock()
	defer bot.mux.RUnlock()
	value, ok := bot.WaitingFrames[key]
	return value, ok
}

func (bot *Bot) delWaitingFrame(key string) {
	bot.mux.Lock()
	defer bot.mux.Unlock()
	delete(bot.WaitingFrames, key)
}

func NewPush(frame *onebot.Frame) error {
	data, err := json.Marshal(frame)
	if err != nil {
		return err
	}
	for _, bot := range Bots {
		bot.Session.Send(websocket.TextMessage, data)
	}
	return nil
}
