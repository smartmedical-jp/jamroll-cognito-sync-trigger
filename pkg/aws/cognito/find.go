package cognito

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"jam-roll-cognito-sync-trigger/pkg/aws/setting"
)

func FindByEmail(email string) (*cognitoidentityprovider.UserType, error) {
	client, err := NewIdpClient()
	if err != nil {
		return nil, err
	}

	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(setting.GetUserPoolID()),
		Filter:     aws.String(fmt.Sprintf("email = \"%s\"", email)),
		Limit:      aws.Int64(1),
	}

	result, err := client.ListUsers(input)
	if err != nil {
		return nil, err
	}

	if len(result.Users) == 0 {
		return nil, nil
	}

	return result.Users[0], nil
}
