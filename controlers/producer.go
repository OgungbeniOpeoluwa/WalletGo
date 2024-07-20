package controlers

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

func Producer(client pulsar.Client, message string) {
	producers, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:  "golang3",
		Schema: pulsar.NewStringSchema(nil),
	})
	if err != nil {
		log.Println("error creating producer ->", err)
		return
	}
	defer producers.Close()

	send, err := producers.Send(context.Background(), &pulsar.ProducerMessage{
		Value: message,
	})
	if err != nil {
		return
	}
	log.Println("successfully sent message", send)

}
