package mq

import "context"

type Msg struct {
	Topic string
	Data  interface{}
}
type Mqer interface {
	Publish(ctx context.Context, topic string, val interface{}) error
	Subscribe(ctx context.Context, topics ...string) chan *Msg
}


