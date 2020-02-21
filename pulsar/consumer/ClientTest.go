package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
		OperationTimeoutSeconds: 5,
		MessageListenerThreads:  runtime.NumCPU(),
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client : %v", err)
	}

	msgChannel := make(chan pulsar.ConsumerMessage)

	consumerOpts := pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "my-tpic-go",
		Type:             pulsar.Exclusive,
		MessageChannel:   msgChannel,
	}

	consumer, err := client.Subscribe(consumerOpts)

	if err != nil {
		log.Fatalf("Counld not establish subscription:%v", err)
	}

	defer consumer.Close()

	for cm := range msgChannel {
		msg := cm.Message
		fmt.Printf("Message ID: %s\n", msg.ID())
		fmt.Printf("Message Value: %s\n", string(msg.Payload()))
		consumer.Ack(msg)
	}

}
