package redis

import (
	"context"

	red "github.com/go-redis/redis/v8"
	"github.com/liujunren93/share_utils/common/mq"
	re "github.com/liujunren93/share_utils/db/redis"
)

type redisMq struct {
	redis *re.Client
	opt   *option
}
type option struct {
	MsgSize int
}

func NewMq(client *re.Client, opts ...func(*option)) *redisMq {
	var opt option
	for _, f := range opts {
		f(&opt)
	}
	if opt.MsgSize == 0 {
		opt.MsgSize = 100
	}
	return &redisMq{redis: client, opt: &opt}
}

func (r *redisMq) UpdateClient(client *re.Client) {
	r.redis = client
}
func (r *redisMq) Publish(ctx context.Context, topic string, val interface{}) error {
	publish := r.redis.Publish(ctx, topic, val)
	return publish.Err()
}

func (r *redisMq) Subscribe(ctx context.Context, topics ...string) (ch chan *mq.Msg) {
	res := r.redis.Subscribe(ctx, topics...)
	ch = make(chan *mq.Msg, r.opt.MsgSize)
	go func() {
		for {
			select {
			case msg := <-res.Channel(red.WithChannelSize(r.opt.MsgSize)):
				ch <- &mq.Msg{
					Topic: msg.Channel,
					Data:  msg.Payload,
				}
			case <-ctx.Done():
				return
			}

		}
	}()
	return ch
}
func WithMsgSIze(size int) func(*option) {
	return func(o *option) {
		o.MsgSize = size
	}
}
