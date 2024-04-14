package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"send-wave/internal/core/services"
	"send-wave/internal/externals/mail"
	"send-wave/internal/handlers"
)

func main() {
	handler := handlers.NewMailHandler(
		services.NewMailService(
			mail.NewLocalMail(),
		),
	)
	lambda.Start(handler.HandleMail)
}
