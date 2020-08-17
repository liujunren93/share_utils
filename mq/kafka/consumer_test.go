package kafka

import (
	"fmt"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	c, _ := NewConsumer(nil, nil)
	p, _ := c.ConsumePartition("micro", 0, 1)
	defer p.Close()

	for {
		select {
		case msg := <-p.Messages():

			fmt.Printf("%s", msg.Value)
		case e := <-p.Errors():
			fmt.Printf("%s", e)

		}
	}

}
