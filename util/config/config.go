package config

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

var (
	DatabaseUsername                 string
	DatabasePassword                 string
	DatabaseHost                     string
	DatabasePort                     string
	DatabaseName                     string
	PaystackSecretKey                string
	PaystackInitializeTransactionUrl string
	ContractCode                     string
	MonifySecretKey                  string
	MonifyApiKey                     string
	MonifyInitializeTransactionUrl   string
)

func Load(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mapConfiguration()
}

func mapConfiguration() {
	mapDatabaseConfiguration()
	mapPaystackConfiguration()
	mapMonifyConfiguration()

}

func mapDatabaseConfiguration() {
	DatabaseUsername = os.Getenv("DB_USERNAME")
	DatabasePassword = os.Getenv("DB_PASSWORD")
	DatabaseHost = os.Getenv("DB_HOST")
	DatabasePort = os.Getenv("DB_PORT")
	DatabaseName = os.Getenv("DB_NAME")

}

func mapPaystackConfiguration() {
	PaystackSecretKey = os.Getenv("PAYSTACK_SECRET_KEY")
	PaystackInitializeTransactionUrl = os.Getenv("PAYSTACK_INITIALIZE_TRANSACTION_URL")
}

func mapMonifyConfiguration() {
	ContractCode = os.Getenv("CONTRACT_CODE")
	MonifyInitializeTransactionUrl = os.Getenv("MONNIFY_INITIALIZE_TRANSACTION")
	MonifyApiKey = os.Getenv("MONNIFY_API_KEY")
	MonifySecretKey = os.Getenv("MONNIFY_SECRET_KEY")
}
