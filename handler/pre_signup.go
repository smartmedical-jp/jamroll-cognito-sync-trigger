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
	//email := event.Request.UserAttributes["email"]

	// log
	err := log.PrintEventLog(event)
	if err != nil {
		return event, err
	}

	//switch event.TriggerSource {
	//case TriggerSourceSignUp, TriggerSourceAdminCreateUser:
	//	exist, err := isUserAlreadyExist(ctx, email)
	//	if err != nil {
	//		return event, err
	//	}
	//	if exist {
	//		return event, ErrUserAlreadyExist
	//	}

	// ソーシャルログインが成功し、そのユーザが Cognito ユーザプールに存在しない場合にトリガされる
	// - ただし、メールアドレス認証とソーシャルログインのようにログイン方式が異なると同じメールアドレスでも異なるユーザとして扱われてしまう
	//	 = Cognito ユーザプールに存在する場合も存在しない場合も呼び出される
	//case TriggerSourceExternalProvider:
	//	// サインイン時
	//	exist, err := firebase.ExistByEmail(ctx, email)
	//	if err != nil {
	//		return event, err
	//	}
	//	if !exist {
	//		return event, firebase.ErrUserNotExist
	//	}
	//	exist, err = cognito.ExistByEmail(email)
	//	if err != nil {
	//		return event, err
	//	}
	//	if exist {
	//		return event, ErrUserAlreadyExist
	//	}
	//
	//	// サインアップ時
	//	exist, err = isUserAlreadyExist(ctx, email)
	//	if err != nil {
	//		return event, err
	//	}
	//	if exist {
	//		return event, ErrUserAlreadyExist
	//	}
	//}

	return event, nil
}

func isUserAlreadyExist(ctx context.Context, email string) (bool, error) {
	exist, err := firebase.ExistByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}

	exist, err = cognito.ExistByEmail(email)
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}

	return false, nil
}
