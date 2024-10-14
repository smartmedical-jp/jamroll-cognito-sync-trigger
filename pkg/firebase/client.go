package firebase

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/service/ssm"
	"google.golang.org/api/option"
	"jam-roll-cognito-sync-trigger/env"
	ssm2 "jam-roll-cognito-sync-trigger/pkg/aws/ssm"

	fb "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func NewClient(ctx context.Context) (*auth.Client, error) {
	firebaseAccessKey, err := getFirebaseAccessKey()
	if err != nil {
		return nil, err
	}

	dec, err := base64.StdEncoding.DecodeString(firebaseAccessKey)
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

func getFirebaseAccessKey() (string, error) {
	ssmClient, err := ssm2.NewClient()
	if err != nil {
		return "", err
	}

	paramKey := fmt.Sprintf("/%s/firebase/access_key", env.GetEnv())
	withDecryption := true
	firebaseAccessKey, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name:           &paramKey,
		WithDecryption: &withDecryption,
	})
	if err != nil {
		return "", err
	}

	return *firebaseAccessKey.Parameter.Value, nil
}
