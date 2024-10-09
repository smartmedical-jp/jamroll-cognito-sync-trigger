//go:generate make -C ../../ gen DIR=migrateuser
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/handler"
)

func main() {
	lambda.Start(handler.MigrateUserHandler)
}
