package redis

import (
	"context"
	"fmt"
	"time"

	re "github.com/go-redis/redis/v8"
)

type Client struct {
	Mode           int8 // general:0 cluster:1 sentinel:2
	Client         *re.Client
	ClusterClient  *re.ClusterClient
	SentinelClient *re.SentinelClient
}
type Configer interface {
	GetMode() int8 // general:0 cluster:
}
type Config struct {
	Mode     int8   // general:0 cluster:1 sentinel:2
	Network  string `json:"network" yaml:"network"`
	Addr     string `json:"addr" yaml:"addr"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}

func (c *Config) GetMode() int8 {
	return c.Mode
}

type ClusterConfig struct {
	Mode     int8     // general:0 cluster:1 sentinel:2
	Addrs    []string `json:"addrs" yaml:"addsr"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
}

func (c *ClusterConfig) GetMode() int8 {
	return c.Mode
}

func NewClient(conf Configer) (cli *Client, err error) {
	cli = new(Client)
	cli.Mode = conf.GetMode()
	switch conf.GetMode() {
	case 0:
		client, err := NewGeneralClient(conf.(*Config))
		if err != nil {
			return nil, err
		}
		cli.Client = client
	case 1:
		client, err := NewClusterClient(conf.(*ClusterConfig))
		if err != nil {
			return nil, err
		}
		cli.ClusterClient = client
	case 2:
		client, err := NewSentinelClient(conf.(*Config))
		if err != nil {
			return nil, err
		}
		cli.SentinelClient = client
	default:
		return nil, fmt.Errorf("mode cannot be supported")
	}
	return

}

func NewGeneralClient(conf *Config) (*re.Client, error) {

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

func NewClusterClient(conf *ClusterConfig) (*re.ClusterClient, error) {

	opt := re.ClusterOptions{
		Addrs:    []string{},
		Username: conf.Username,
		Password: conf.Password,
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	cli := re.NewClusterClient(&opt)

	err := cli.Ping(ctx).Err()
	return cli, err
}

func NewSentinelClient(conf *Config) (*re.SentinelClient, error) {
	opt := re.Options{
		Network:  conf.Network,
		Addr:     conf.Addr,
		Username: conf.Username,
		Password: conf.Password,
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	cli := re.NewSentinelClient(&opt)
	err := cli.Ping(ctx).Err()
	return cli, err
}
