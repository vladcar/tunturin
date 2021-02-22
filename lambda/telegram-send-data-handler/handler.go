package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"tunturin/services/telegram-service"
)

func Handle(event events.APIGatewayV2HTTPRequest) error {
	_, er :=telegram_service.SendMessage(event.Body)
	return er
}

func main() {
	lambda.Start(Handle)
}
