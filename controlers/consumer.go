package controlers

import (
	"WalletService/dtos/request"
	"WalletService/services"
	"context"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

func Consumer(client pulsar.Client) {
	jsonSchemaDef := `{
  "type": "record",
  "name": "CreateAccountRequest",
  "namespace": "com.example.pulsar.testing",
  "fields": [
    {
      "name": "email",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "firstName",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "lastName",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "password",
      "type": ["null", "string"],
      "default": null
    },
    {
      "name": "phoneNumber",
      "type": ["null", "string"],
      "default": null
    }
  ]
}`
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		// fill `Topic` field will create a single-topic consumer
		Topic:            "account2",
		SubscriptionName: "my-subscription",
		Type:             pulsar.Exclusive,
		Schema:           pulsar.NewJSONSchema(jsonSchemaDef, nil),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Println("Failed to receive message:", err)
			continue
		}
		log.Printf("Received message msgId: %v -- content: '%s'", msg.ID(), string(msg.Payload()))
		var userRequest request.CreateAccountRequest
		err = json.Unmarshal(msg.Payload(), &userRequest)
		if err != nil {
			log.Println("Failed to unmarshal CreateUserRequest:", err)
		}
		user, err := services.NewWalletServiceImpl().CreateAccount(&userRequest)
		if err != nil {
			return
		}
		Producer(client, user)
		log.Println("Processed and sending user info")

		err = consumer.Ack(msg)
		if err != nil {
			return
		}
	}
}
