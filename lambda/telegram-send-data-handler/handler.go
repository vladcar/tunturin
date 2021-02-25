package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"tunturin/services/telegram-webhook"
)

func Handle(event events.APIGatewayV2HTTPRequest) error {
	er := telegram_webhook.HandleWebhook(event.Body)
	if er != nil {
		log.Fatal(er)
	}
	return nil
}

func main() {
	lambda.Start(Handle)
}
