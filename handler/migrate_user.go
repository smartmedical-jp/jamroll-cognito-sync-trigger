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
	TriggerSourceForgotPassword = "UserMigration_ForgotPassword"
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
		event, err = createUserInUserPool(event)
		if err != nil {
			return event, err
		}
	// パスワードリセットがリクエストされ、そのリクエストユーザがユーザプールに存在しない場合のみ、このトリガーが発生する
	// NOTE:
	// - リクエスト時点では、パスワードリセット対象メールアドレスだけ送信される
	// - 新パスワードを設定するためのUIは別途任意のタイミングで表示必要がある
	case TriggerSourceForgotPassword:
		// TODO 1. Firebase に存在するか確認
		// TODO 2. Firebase に存在する場合、新パスワードを設定するためのUIを表示
		// 存在しない場合は、エラーを返す？
		// TODO 3. Firebase でパスワードリセット処理を行う（何かあって Firebase にまた戻す、みたいな状況を考慮（API経由だけで無理かも/要らんかも））
		// TODO 4. Firebase でパスワードリセット処理が成功したら、ユーザープールに メールアドレス + 新パスワード でユーザを作成
	}

	return event, nil
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
