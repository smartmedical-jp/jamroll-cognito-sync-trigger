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
	event.CognitoEventUserPoolsMigrateUserResponse.UserAttributes = map[string]string{
		"email":          event.UserName,
		"email_verified": "true",
	}
	event.CognitoEventUserPoolsMigrateUserResponse.FinalUserStatus = "CONFIRMED"
	event.CognitoEventUserPoolsMigrateUserResponse.MessageAction = "SUPPRESS"

	return event, nil
}

// Firebase にユーザが存在するか確認
//func checkUserExistInFirebase(event events.CognitoEventUserPoolsMigrateUser) error {
//	// Firebase にユーザが存在するか確認
//
//	return nil
//}
