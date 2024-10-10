//go:generate make -C ../../ gen DIR=presignup
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"jam-roll-cognito-sync-trigger/handler"
)

func main() {
	lambda.Start(handler.PreSignupHandler)
}
