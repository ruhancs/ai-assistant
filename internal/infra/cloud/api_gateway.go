package cloud

import (
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ruhancs/ai-assistant/internal/usecase"
)

func Proccess(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//request tera dados do twillio,
	//result tera a pergunta recebida pelo twillio, whatsapp
	result, err := parseBase64RequestData(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	//enviar a pergunta para o chat-gpt e obter a resposta
	text, err := usecase.GenerateGPTText(result)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       text,
	}, nil
}

// pegar os dados do twillio convertido para utilizar na api gateway
func parseBase64RequestData(req string) (string, error) {
	//pegar os dados base64 e converter para texto puro
	dataBytes, err := base64.StdEncoding.DecodeString(req)
	if err != nil {
		return "", err
	}

	//pegar campo por campo da query string
	//dados do twillio
	data, err := url.ParseQuery(string(dataBytes))
	if data.Has("Body") {
		return data.Get("Body"), nil
	}
	return "", errors.New("body not found")
}
