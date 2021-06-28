package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://172.16.60.103:6650",
	})
	if err != nil {
		return
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-delay-topic",
	})
	if err != nil {
		return
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload:      []byte("hello"),
		DeliverAfter: 30 * time.Second,
	})
	if err != nil {
		return
	}
	defer producer.Close()

	fmt.Println("Published delay message")
}
