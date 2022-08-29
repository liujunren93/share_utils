package redis

import (
	"context"
	"time"

	re "github.com/go-redis/redis/v8"
)

func NewRedis(conf *Config) (*re.Client, error) {

	opt := re.Options{
		Network:  conf.Network,
		Addr:     conf.Addr,
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	cli := re.NewClient(&opt)
	err := cli.Ping(ctx).Err()
	return cli, err
}

type Config struct {
	Network  string `js`
	Addr     string
	Username string
	Password string
	DB       int
}
