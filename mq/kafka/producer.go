package kafka

import "github.com/Shopify/sarama"

func defaultProducerConfig() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Idempotent = true
	conf.Net.MaxOpenRequests = 1
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	conf.Version = sarama.V2_5_0_0
	conf.Producer.RequiredAcks=-1
	return conf
}


func NewAsyncProducer(address []string, config *sarama.Config) (sarama.AsyncProducer, error) {
	if config == nil {
		config = defaultProducerConfig()
	}
	if address == nil {
		address = []string{"localhost:2182"}
	}
	return sarama.NewAsyncProducer(address, config)
}

func NewSyncProducer(address []string, config *sarama.Config) (sarama.SyncProducer, error) {
	if config == nil {
		config = defaultProducerConfig()
	}
	if address == nil {
		address = []string{"localhost:9092"}
	}
	return sarama.NewSyncProducer(address, config)
}
