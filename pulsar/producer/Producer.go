package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

func main() {

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})

	if err != nil {
		log.Fatal(err)
	}

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		msg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("message-%d", i)),
		}
		if err := producer.Send(ctx, msg); err != nil {
			log.Fatal(err)
		}
		asyncMsg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("async-message-%d", i)),
		}
		producer.SendAsync(ctx, asyncMsg, func(msg pulsar.ProducerMessage, err error) {
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("the %s successfull published", string(msg.Payload))
		})
	}

}
