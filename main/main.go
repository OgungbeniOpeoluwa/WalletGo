package main

import (
	"WalletService/db"
	"WalletService/logger"
	"WalletService/router"
	"WalletService/util/config"
	"context"
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
	err = r.Run(":9088")
	if err != nil {
		return
	}

}
