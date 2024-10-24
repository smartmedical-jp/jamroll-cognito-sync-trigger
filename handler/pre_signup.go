package handler

import (
	"context"
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
		exist, _ := firebase.ExistByEmail(ctx, email)
		if exist {
			return event, firebase.ErrUserAlreadyExist
		}
		exist, err := cognito.ExistByEmail(email)
		if err != nil {
			return event, err
		}
		if exist {
			return event, cognito.ErrUserAlreadyExist
		}

	/*
		・移行トリガーがトリガされるか、
		・直接CDKなどでこの操作を行う
		場合にトリガされる

		移行トリガーはユーザプールにログインユーザが存在しない場合にのみトリガされるので、既存ユーザとの重複チェックは不要
		また、SOLID にならないが、migrate_user コントローラー（移行トリガー発火時の実装）で firebase に存在するかチェックしているためここではチェックしない
		（※ 米国のデータストアに毎回確認しにいく firebase 処理を行う回数を減らすことを優先）

		さらに、現状（2024/10/09）機能開発のなかでそれが使用される可能性は低いので後者も無視
	*/
	case TriggerSourceAdminCreateUser:

	/*
		ソーシャルログインが成功し、そのユーザが Cognito ユーザプールに存在しない場合に、そのユーザを登録しようとしてトリガされる

		JamRoll では以下のケースでトリガされる
		1. 移行未完了ユーザのソーシャルログイン
		   - firebase にはユーザが存在し、Cognito にユーザが存在しない
		2. サインアップ時のソーシャルログイン
		   - firebase にも Cognito にもユーザが存在しない

		移行未完了ユーザのソーシャルログインの場合は、firebase にユーザが存在する必要がある。
		逆に、サインアップの場合は firebase に存在しないユーザである必要がある。
		```
		// 移行未完了ユーザのソーシャルログイン
		exist, err := firebase.ExistByEmail(ctx, email)
		if err != nil || !exist {
			return event, err
		}
		// サインアップの場合
		exist, _ = firebase.ExistByEmail(ctx, email)
		if exist {
			return event, firebase.ErrUserAlreadyExist
		}
		// その他の処理...
		```
		移行未完了ユーザのソーシャルログイン or サインアップにおいて、Cognito ユーザプールに存在しない場合毎回このトリガーが発火されるが、
		エラーにすべきケースが真逆であるため、移行未完了ユーザのソーシャルログインなのかサインアップなのかを区別する仕組みが必要。しかし現状 Cognito にはその仕組みがない。
		=> よって、それはフロントで制御するしかない
	*/
	case TriggerSourceExternalProvider:
		// 認証方法が違うと、メールアドレスが同じでも異なるユーザとして扱われてしまうため、重複するユーザが存在する場合統合する
		// ※ このトリガが発火する時点でログイン自体は問題なく成功しているので、エラーにはせず、ユーザ登録のみブロックしてアプリケーションにリダイレクトすることになる
		// https://qiita.com/Naoki1126/items/e6294f0ed189344a5bd7
		sameUser, err := cognito.FindByEmail(email)
		if err != nil {
			return event, err
		}
		if sameUser != nil {
			provider := cognito.GetExternalProvider(event)
			err := cognito.AdminLinkUser(sameUser, event, provider)
			if err != nil {
				return event, err
			}
			return event, nil
		}
		event.Request.UserAttributes["email_verified"] = "true"
	}

	return event, nil
}
