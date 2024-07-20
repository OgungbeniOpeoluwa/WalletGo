package main

import (
	"WalletService/controlers"
	"WalletService/db"
	"WalletService/logger"
	"WalletService/router"
	"WalletService/util/config"
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	config.Load("../.env")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.DbConnection()
	if err != nil {
		logger.ErrorLogger(err)
		return
	}
	//DB = dbs
	r := gin.Default()
	router.Routes(r)
	log.Println("connected")
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}
	controlers.Consumer(client)
	defer client.Close()
	err = r.Run(":9088")
	if err != nil {
		return
	}

}
