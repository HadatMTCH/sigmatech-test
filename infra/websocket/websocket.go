package websocket

import (
	"base-api/config"
	"net/url"
	"os"
	"os/signal"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

type websckt struct {
	cfg *config.ServerConfig
}

type WebsocketInterface interface {
	EmitEvent(path string, data []byte)
}

func NewWebsocket(cfg *config.ServerConfig) WebsocketInterface {
	return &websckt{
		cfg: cfg,
	}
}

func (w websckt) EmitEvent(path string, data []byte) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	websocketAddr := strings.Split(w.cfg.WebsocketAddr, "://")
	u := url.URL{Scheme: websocketAddr[0], Host: websocketAddr[1], Path: path}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Error("dial:", err)
		return
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Println("write:", err)
		return
	}
}
