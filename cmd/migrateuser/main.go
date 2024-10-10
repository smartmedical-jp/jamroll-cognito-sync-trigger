//go:generate make -C ../../ gen DIR=migrateuser
package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/handler"
	"github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/pkg/registry"
	"log"
)

func main() {
	ctx := context.Background()

	rgst, err := registry.NewRegistry(ctx)
	if err != nil {
		log.Fatalf("failed to initialize registry: %v", err)
	}
	handler := handler.NewHandler(*rgst)

	lambda.Start(handler.MigrateUserHandler)
}
