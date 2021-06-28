package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

// https://github.com/apache/pulsar-client-go

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://172.16.60.103:6650", //更换为接入点地址
		// ListenerName:      "custom:1300*0/vpc-**/subnet-****",        //更换为路由 ID
		// Authentication:    pulsar.NewAuthenticationToken("eyJh****"), //更换为密钥
		//OperationTimeout:  30 * time.Second,
		//ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}
	defer producer.Close()

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})
	if err != nil {
		fmt.Println("Failed to publish message", err)
	}

	fmt.Println("Published message")
}
