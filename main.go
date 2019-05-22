package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// request.RequestContext
	// switch request.HTTPMethod	{
	// case "GET"
	// 	if (request.Path == "")
	// case "POST"
	// case "DELETE"
	// }

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(request.Body),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
