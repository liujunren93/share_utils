package context

import (
	"sync"
	"time"
)

type ShContext struct {
	Header *sync.Map
}

func NewContext() *ShContext {
	return &ShContext{
		Header: &sync.Map{},
	}
}

func (*ShContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*ShContext) Done() <-chan struct{} {
	return nil
}

func (*ShContext) Err() error {
	return nil
}

func (*ShContext) Value(key interface{}) interface{} {
	return nil
}
