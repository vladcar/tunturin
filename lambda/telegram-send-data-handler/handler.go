package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"tunturin/services/telegram-webhook"
)

func Handle(event events.APIGatewayV2HTTPRequest) error {
	fmt.Println(event.Body)
	if er := telegram_webhook.HandleWebhook(event.Body); er != nil {
		log.Fatal(er)
	}
	return nil
}

func main() {
	lambda.Start(Handle)
}
