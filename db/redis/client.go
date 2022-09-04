package redis

import (
	"context"
	"fmt"
	"time"

	re "github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
)

type Client struct {
	Mode int8 // general:0 cluster:1 sentinel:2
	Cmdable
}
type Configer interface {
	GetMode() int8 // general:0 cluster:
	SetMode(int8)
}

type Base struct {
	Mode int8 `json:"mode" yaml:"mode"` // general:1 cluster:2 sentinel:3
}

func (c *Base) GetMode() int8 {
	return c.Mode
}
func (c *Base) SetMode(mode int8) {
	c.Mode = mode
}

type Config struct {
	Base
	Network  string `json:"network" yaml:"network"`
	Addr     string `json:"addr" yaml:"addr"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}

type ClusterConfig struct {
	Base
	Addrs    []string `json:"addrs" yaml:"addsr"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
}

type SentinelConfig struct {
	Base
	DB               int8     `json:"db" yaml:"db"`
	MasterName       string   `json:"masterName" yaml:"masterName"`
	Addrs            []string `json:"addrs" yaml:"addsr"`
	SentinelUsername string   `json:"sentinelUsername" yaml:"sentinelUsername"`
	SentinelPassword string   `json:"sentinelPassword" yaml:"sentinelPassword"`
	Username         string   `json:"username" yaml:"username"`
	Password         string   `json:"password" yaml:"password"`
}

func NewClient(redisConf map[string]interface{}) (*Client, error) {
	var base Base
	var conf Configer
	err := mapstructure.Decode(redisConf, &base)
	if err != nil {
		return nil, err
	}
	if base.GetMode() == 1 {
		conf = &Config{}
	} else if base.GetMode() == 2 {
		conf = &ClusterConfig{}
	} else if base.GetMode() == 3 {
		conf = &SentinelConfig{}
	} else {
		panic("ConfCenter redis config mod err ")
	}
	conf.SetMode(base.GetMode())
	mapstructure.Decode(redisConf, conf)
	var cli = new(Client)
	cmd, err := newClient(conf)
	if err != nil {
		return nil, err
	}
	cli.Mode = conf.GetMode()
	cli.Cmdable = cmd
	return cli, nil
}
func newClient(conf Configer) (Cmdable, error) {
	switch conf.GetMode() {
	case 1:
		return newGeneralClient(conf.(*Config))

	case 2:
		return newClusterClient(conf.(*ClusterConfig))

	case 3:
		return newSentinelClient(conf.(*SentinelConfig))

	default:
		return nil, fmt.Errorf("mode cannot be supported")
	}

}

func newGeneralClient(conf *Config) (*re.Client, error) {

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

func newClusterClient(conf *ClusterConfig) (*re.ClusterClient, error) {

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

func newSentinelClient(conf *SentinelConfig) (*re.Client, error) {
	opt := re.FailoverOptions{
		MasterName:       conf.MasterName,
		SentinelAddrs:    conf.Addrs,
		SentinelUsername: conf.SentinelUsername,
		SentinelPassword: conf.SentinelPassword,

		Username: conf.Username,
		Password: conf.Password,
		DB:       int(conf.DB),
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	cli := re.NewFailoverClient(&opt)
	err := cli.Ping(ctx).Err()
	return cli, err

}
