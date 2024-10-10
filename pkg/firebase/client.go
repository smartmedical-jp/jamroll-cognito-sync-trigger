package firebase

import (
	"context"
	"encoding/base64"
	"fmt"
	ssm2 "github.com/aws/aws-sdk-go/service/ssm"
	"github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/pkg/aws/ssm"
	"google.golang.org/api/option"

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
	ssmClient, err := ssm.NewClient()
	if err != nil {
		return "", err
	}
	paramKey := "/dev/firebase/access_key"
	withDecryption := true

	firebaseAccessKey, err := ssmClient.GetParameter(&ssm2.GetParameterInput{
		Name:           &paramKey,
		WithDecryption: &withDecryption,
	})
	if err != nil {
		fmt.Println("failed to get firebase access key", err)
		return "", err
	}

	return *firebaseAccessKey.Parameter.Value, nil
}
