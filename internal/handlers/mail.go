package handlers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log/slog"
	"net/http"
	"send-wave/internal/core/ports"
	"send-wave/internal/dto"
	"time"
)

type MailHandler struct {
	service ports.MailService
}

func NewMailHandler(service ports.MailService) *MailHandler {
	return &MailHandler{
		service: service,
	}
}

func (m MailHandler) HandleMail(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := decodeAndValidateBody(request)
	if err != nil {
		return createDefaultJSONResponse(
			http.StatusBadRequest,
			"Not a valid request",
		)
	}

	startTime := time.Now()

	err = m.service.SendMail(req.Message)
	if err != nil {
		return createDefaultJSONResponse(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	timeElapsed := time.Since(startTime)

	response := dto.MailResponse{
		Details:     map[string]int{"client_1": 10},
		ElapsedTime: timeElapsed.String(),
	}

	return createDefaultJSONResponse(
		http.StatusOK,
		response,
	)
}

func decodeAndValidateBody(request events.APIGatewayProxyRequest) (*dto.MailMessage, error) {
	var req *dto.MailMessage
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}
	return req, nil
}

func createDefaultJSONResponse(
	statusCode int,
	body interface{},
) (events.APIGatewayProxyResponse, error) {
	binaryBody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	slog.Info("mail sender response:", map[string]any{
		"status_code": statusCode,
		"body":        string(binaryBody),
	})

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(binaryBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
