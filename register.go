package bot

import (
	"github.com/2mf8/Better-Bot-Go/log"
	"github.com/2mf8/Better-Bot-Go/openapi"
	"github.com/2mf8/Better-Bot-Go/websocket2"
)

// SetLogger 设置 logger，需要实现 sdk 的 log.Logger 接口
func SetLogger(logger log.Logger) {
	log.DefaultLogger = logger
}

// SetSessionManager 注册自己实现的 session manager
func SetSessionManager(m SessionManager) {
	defaultSessionManager = m
}

// SetWebsocketClient 替换 websocket 实现
func SetWebsocketClient(c websocket2.WebSocket) {
	websocket2.Register(c)
}

// SetOpenAPIClient 注册 openapi 的不同实现，需要设置版本
func SetOpenAPIClient(v openapi.APIVersion, c openapi.OpenAPI) {
	openapi.Register(v, c)
}
