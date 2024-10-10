package firebase

import (
	"context"
	"encoding/base64"
	"google.golang.org/api/option"
	"os"

	fb "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func NewClient(ctx context.Context) (*auth.Client, error) {
	// TODO 環境ごとのアクセスキーを取得する
	dec, err := base64.StdEncoding.DecodeString(os.Getenv("FIREBASE_ACCESS_KEY"))
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(dec)
	app, err := fb.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
