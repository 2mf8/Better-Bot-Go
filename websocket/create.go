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
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	NewBot(xBotSelfId, c)
	return nil
}
