package cognito

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"jam-roll-cognito-sync-trigger/pkg/aws/setting"
)

var (
	ErrUserAlreadyExist = errors.New("cognito user pool: user already exists")
)

func ExistByEmail(email string) (bool, error) {
	client, err := NewIdpClient()
	if err != nil {
		return false, err
	}

	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(setting.GetUserPoolID()),
		Filter:     aws.String(fmt.Sprintf("email = \"%s\"", email)),
		Limit:      aws.Int64(1),
	}

	result, err := client.ListUsers(input)
	if err != nil {
		return false, err
	}

	if len(result.Users) == 0 {
		return false, nil
	}

	return true, nil
}
