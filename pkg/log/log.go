package log

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

// PrintEventLog イベント全体をログに出力
func PrintEventLog[
	T events.CognitoEventUserPoolsMigrateUser | events.CognitoEventUserPoolsPreSignup,
](event T) error {
	eventJSON, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		log.Printf("Error marshaling event: %v", err)
		return err
	}
	log.Println("Received Cognito event:")
	log.Println(string(eventJSON))

	return nil
}
