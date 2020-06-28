package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Printf("EVENT: %s", request.Body)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello world",
	}, nil
}

func main() {
	lambda.Start(handler)
}
