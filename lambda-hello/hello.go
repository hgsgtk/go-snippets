package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"What is your name?"`
}

type Response struct {
	Message string `json:"Answer:"`
}

func hello(event Event) (Response, error) {
	return Response{Message: fmt.Sprintf("Hello %s!!", event.Name)}, nil
}

func main() {
	lambda.Start(hello)
}
