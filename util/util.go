package util

import (
	"runtime/debug"
	"strconv"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

var GlobalId int64 = 0

func SafeGo(fn func()) {
	go func() {
		defer func() {
			e := recover()
			if e != nil {
				log.Errorf("err recovered: %+v", e)
				log.Errorf("%s", debug.Stack())
			}
		}()
		fn()
	}()
}

func GenerateId() int64 {
	return atomic.AddInt64(&GlobalId, 1)
}

func GenerateIdStr() string {
	return strconv.FormatInt(GenerateId(), 10)
}