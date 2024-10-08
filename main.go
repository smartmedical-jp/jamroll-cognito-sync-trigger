package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsMigrateUser) (events.CognitoEventUserPoolsMigrateUser, error) {
	// イベント全体をログに出力
	eventJSON, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		log.Printf("Error marshaling event: %v", err)
		return event, err
	}

	log.Println("Received Cognito event:")
	log.Println(string(eventJSON))

	// 他のロジックを実装する場合、ここに処理を追加
	return event, nil
}

func main() {
	lambda.Start(handler)
}
