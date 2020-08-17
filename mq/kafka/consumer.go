package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
)

func defaultConsumerConfig() *sarama.Config {
	conf := sarama.NewConfig()

	conf.Consumer.Return.Errors = true
	conf.Version = sarama.V2_5_0_0
	return conf

}

// NewConsumer return Consumer
func NewConsumer(address []string, conf *sarama.Config) (sarama.Consumer, error) {
	if address == nil {
		address = []string{"localhost:9092"}
	}
	if conf != nil {
		conf = defaultConsumerConfig()
	}

	return sarama.NewConsumer(address, conf)
}

// NewConsumer return ConsumerGroup
func NewConsumerGroup(address []string, group string, conf *sarama.Config) (sarama.ConsumerGroup, error) {
	if address == nil {
		address = []string{"localhost:9092"}
	}
	if conf != nil {
		conf = defaultConsumerConfig()
	}
	if group == "" {
		return nil, errors.New("group cannot be empty")
	}

	return sarama.NewConsumerGroup(address, group, conf)

}
