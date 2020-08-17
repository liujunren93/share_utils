package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
)

func TestNewAsyncProducer(t *testing.T) {
	producer, err := NewSyncProducer(nil,nil)

	message, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: "micro",
		Value: sarama.StringEncoder("new"),
		Partition: 0,
	})
fmt.Println(message, offset, err)
}
