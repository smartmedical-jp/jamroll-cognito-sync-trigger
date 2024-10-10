package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"jam-roll-cognito-sync-trigger/pkg/log"
)

func PreSignupHandler(event events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	err := log.PrintEventLog(event)
	if err != nil {
		return event, err
	}

	return event, nil
}
