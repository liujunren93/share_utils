package redis

import (
	"context"
	"time"

	re "github.com/go-redis/redis/v8"
)

func NewRedis(conf *re.Options) (*re.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	cli := re.NewClient(conf)
	err := cli.Ping(ctx).Err()
	return cli, err
}
