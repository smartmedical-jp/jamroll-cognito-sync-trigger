package handler

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"jam-roll-cognito-sync-trigger/pkg/aws/cognito"
	"jam-roll-cognito-sync-trigger/pkg/aws/setting"
	"jam-roll-cognito-sync-trigger/pkg/firebase"
	"jam-roll-cognito-sync-trigger/pkg/log"
)

const (
	// https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/cognito-user-pools-working-with-lambda-triggers.html#working-with-lambda-trigger-sources
	TriggerSourceSignUp           = "PreSignUp_SignUp"
	TriggerSourceAdminCreateUser  = "PreSignUp_AdminCreateUser"
	TriggerSourceExternalProvider = "PreSignUp_ExternalProvider"
)

var (
	ErrUserAlreadyExist = errors.New("cognito user pool: user already exists")
)

func PreSignupHandler(
	ctx context.Context,
	event events.CognitoEventUserPoolsPreSignup,
) (events.CognitoEventUserPoolsPreSignup, error) {
	// 初期設定
	setting.InitSetting(setting.PreSignup{Event: event})
	email := event.Request.UserAttributes["email"]

	// log
	err := log.PrintEventLog(event)
	if err != nil {
		return event, err
	}

	switch event.TriggerSource {
	case TriggerSourceSignUp:
		exist, _ := firebase.ExistByEmail(ctx, event.UserName)
		if exist {
			return event, firebase.ErrUserAlreadyExist
		}
		exist, err := cognito.ExistByEmail(email)
		if err != nil {
			return event, err
		}
		if exist {
			return event, ErrUserAlreadyExist
		}
	case TriggerSourceAdminCreateUser:
		exist, err := firebase.ExistByEmail(ctx, event.UserName)
		if err != nil || !exist {
			return event, err
		}
		exist, err = cognito.ExistByEmail(email)
		if err != nil {
			return event, err
		}
		if exist {
			return event, ErrUserAlreadyExist
		}
	// ソーシャルログインが成功し、そのユーザが Cognito ユーザプールに存在しない場合にトリガされる
	// - ただし、メールアドレス認証とソーシャルログインのようにログイン方式が異なると同じメールアドレスでも異なるユーザとして扱われてしまう
	//	 = Cognito ユーザプールに存在する場合も存在しない場合も呼び出される
	case TriggerSourceExternalProvider:
		// TODO 以下はサインイン時のみ実行したい
		exist, err := firebase.ExistByEmail(ctx, email)
		if err != nil {
			return event, err
		}
		if !exist {
			return event, firebase.ErrUserNotExist
		}

		// サインアップの場合
		exist, _ = firebase.ExistByEmail(ctx, event.UserName)
		if exist {
			return event, firebase.ErrUserAlreadyExist
		}
		exist, err = cognito.ExistByEmail(email)
		if err != nil {
			return event, err
		}
		if exist {
			return event, ErrUserAlreadyExist
		}
	}

	return event, nil
}
