package handler

import (
	"github.com/aws/aws-lambda-go/events"
	log2 "github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/pkg/log"
)

func PreSignupHandler(event events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	err := log2.PrintEventLog(event)
	if err != nil {
		return event, err
	}

	return event, nil
}
