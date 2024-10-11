package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"jam-roll-cognito-sync-trigger/pkg/firebase"
	"jam-roll-cognito-sync-trigger/pkg/log"
)

const (
	// https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/user-pool-lambda-migrate-user.html#user-pool-lambda-migrate-user-trigger-source
	TriggerSourceAuthentication = "UserMigration_Authentication"
)

func MigrateUserHandler(
	ctx context.Context,
	event events.CognitoEventUserPoolsMigrateUser,
) (events.CognitoEventUserPoolsMigrateUser, error) {
	err := log.PrintEventLog(event)
	if err != nil {
		return event, err
	}

	switch event.TriggerSource {
	case TriggerSourceAuthentication:
		err = firebase.ExistByEmail(ctx, event.UserName)
		if err != nil {
			return event, err
		}
		event, err = migrateUser(event)
		if err != nil {
			return event, err
		}
	}

	return event, nil
}

// ユーザープールにユーザーを移行する
func migrateUser(event events.CognitoEventUserPoolsMigrateUser) (events.CognitoEventUserPoolsMigrateUser, error) {
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
