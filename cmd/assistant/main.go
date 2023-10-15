package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ruhancs/ai-assistant/internal/infra/cloud"
)

func main() {
	lambda.Start(cloud.Proccess)
}