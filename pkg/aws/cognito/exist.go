package cognito

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"jam-roll-cognito-sync-trigger/pkg/aws/setting"
)

func ExistByEmail(email string) (bool, error) {
	client, err := NewIdpClient()
	if err != nil {
		return false, err
	}

	user, err := client.AdminGetUser(&cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(setting.GetUserPoolID()),
		Username:   aws.String(email),
	})
	if err != nil {
		fmt.Println("おそらくここが発生されてる", "failed to get user: ", err)
		return false, err
	}
	if user == nil {
		return false, nil
	}

	return true, nil
}
