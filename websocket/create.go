package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UpgradeWebsocket(w http.ResponseWriter, r *http.Request) error {
	xBotSelfId := r.Header.Get("x-bot-self-id")
	addr := r.RemoteAddr
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	NewBot(xBotSelfId, addr, c)
	return nil
}

func UpgradeWebsocketWithSecret(w http.ResponseWriter, r *http.Request) error {
	xBotSelfId := r.Header.Get("x-bot-self-id")
	xBotSecret := r.Header.Get("x-bot-secret")
	addr := r.RemoteAddr
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	NewSecretBot(xBotSelfId, xBotSecret, addr, c)
	return nil
}
