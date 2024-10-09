package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/pkg/log"
)

func MigrateUserHandler(event events.CognitoEventUserPoolsMigrateUser) (events.CognitoEventUserPoolsMigrateUser, error) {
	err := log.PrintEventLog(event)
	if err != nil {
		return event, err
	}

	return event, nil
}
