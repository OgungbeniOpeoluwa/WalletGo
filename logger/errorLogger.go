package logger

import (
	"log"
	"os"
)

func ErrorLogger(err error) {
	logger, errs := os.OpenFile("../errorlogs/error",
		os.O_RDWR|os.O_APPEND|os.O_WRONLY, 0666)
	if errs != nil {
		log.Println(errs)
	}
	_, errs = logger.WriteString(err.Error())
	if errs != nil {
		log.Print(errs)
	}
}
