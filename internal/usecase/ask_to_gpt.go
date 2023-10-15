package usecase

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ruhancs/ai-assistant/internal/domain/entity"
)

func GenerateGPTText(query string) (string, error) {
	req := entity.Request{
		Model: "gpt-3.5-turbo",
		Messages: []entity.Message{
			{
				Role:    "user",
				Content: query,
			},
		},
		MaxTokens: 150,
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	chatGPTUrl := "https://api.openai.com/v1/chat/completions"
	sendToGPT, err := http.NewRequest("POST", chatGPTUrl, bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}
	sendToGPT.Header.Set("Content-Type", "application/json")
	authToken := "Bearer " + os.Getenv("TOKEN_AI")
	sendToGPT.Header.Set("Authorization", authToken)

	response, err := http.DefaultClient.Do(sendToGPT)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	msgFromGPT, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var respFromGPT entity.Response
	err = json.Unmarshal(msgFromGPT, &respFromGPT)
	if err != nil {
		return "", err
	}

	return respFromGPT.Choices[0].Message.Content, nil
}
