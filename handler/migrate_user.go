package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/pkg/log"
)

const (
	// https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/user-pool-lambda-migrate-user.html#user-pool-lambda-migrate-user-trigger-source
	TriggerSourceAuthentication = "UserMigration_Authentication"
	TriggerSourceForgotPassword = "UserMigration_ForgotPassword"
)

func MigrateUserHandler(event events.CognitoEventUserPoolsMigrateUser) (events.CognitoEventUserPoolsMigrateUser, error) {
	err := log.PrintEventLog(event)
	if err != nil {
		return event, err
	}

	switch event.TriggerSource {
	case TriggerSourceAuthentication:
		err = checkUserExistInFirebase(event)
		if err != nil {
			return event, err
		}
		event, err = createUserInUserPool(event)
		if err != nil {
			return event, err
		}
	case TriggerSourceForgotPassword:
	}

	return event, nil
}

// Firebase ログインできるか確認
func checkUserExistInFirebase(event events.CognitoEventUserPoolsMigrateUser) error {
	// TODO Firebase にログインできるか確認する処理を実装
	return nil
}

// ユーザープールにユーザを作成
func createUserInUserPool(event events.CognitoEventUserPoolsMigrateUser) (events.CognitoEventUserPoolsMigrateUser, error) {
	event.CognitoEventUserPoolsMigrateUserResponse.UserAttributes = map[string]string{
		"email":          event.UserName,
		"email_verified": "true",
	}
	// 現状（2024/10/09）、以下2つの設定が必要
	// https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/user-pool-lambda-migrate-user.html#cognito-user-pools-lambda-trigger-syntax-user-migration
	event.CognitoEventUserPoolsMigrateUserResponse.FinalUserStatus = "CONFIRMED"
	event.CognitoEventUserPoolsMigrateUserResponse.MessageAction = "SUPPRESS"

	return event, nil
}
