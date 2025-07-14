package webhook

import (
	"sync"

	"github.com/2mf8/Better-Bot-Go/dto"
)

var eventMu sync.RWMutex

type EventHandler[T any] struct {
	handlers []func(bot *BotHeaderInfo, event *dto.WSPayload, data T)
}

func (eh *EventHandler[T]) Subscribe(handler func(bot *BotHeaderInfo, event *dto.WSPayload, data T)) {
	eventMu.Lock()
	defer eventMu.Unlock()
	newHandlers := make([]func(bot *BotHeaderInfo, event *dto.WSPayload, data T), len(eh.handlers)+1)
	copy(newHandlers, eh.handlers)
	newHandlers[len(eh.handlers)] = handler
	eh.handlers = newHandlers
}
