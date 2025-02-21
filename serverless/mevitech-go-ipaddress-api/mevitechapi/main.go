package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "We did not get your public IP in this request\n"
	} else {
		greeting = fmt.Sprintf("Your public IP is %s!\n", sourceIP)
	}

	responseHeader := map[string]string{
		"Content-Type": "text/plain",
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
		Headers:    responseHeader,
	}, nil
}

func main() {
	lambda.Start(handler)
}
