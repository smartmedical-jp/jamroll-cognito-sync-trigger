package cognito

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"jam-roll-cognito-sync-trigger/pkg/aws/setting"
	"strings"
)

// AdminLinkUser Cognito では認証方法が違うとメールアドレスが同じでも異なるユーザとして扱われてしまうため、統合する
func AdminLinkUser(
	destinationUser *cognitoidentityprovider.UserType,
	event events.CognitoEventUserPoolsPreSignup,
	provider ExternalProvider,
) error {
	client, err := NewIdpClient()
	if err != nil {
		return err
	}

	// https://docs.aws.amazon.com/ja_jp/cognito-user-identity-pools/latest/APIReference/API_AdminLinkProviderForUser.html
	sourceUserName := strings.Split(event.UserName, "_")[0]
	if sourceUserName == "" {
		return fmt.Errorf("sourceUserName is empty")
	}
	res, err := client.AdminLinkProviderForUser(&cognitoidentityprovider.AdminLinkProviderForUserInput{
		DestinationUser: &cognitoidentityprovider.ProviderUserIdentifierType{
			ProviderAttributeValue: aws.String(*destinationUser.Username),
			ProviderName:           aws.String("Cognito"),
		},
		SourceUser: &cognitoidentityprovider.ProviderUserIdentifierType{
			ProviderAttributeName:  aws.String("Cognito_Subject"),
			ProviderAttributeValue: aws.String(sourceUserName),
			ProviderName:           aws.String(provider.String()),
		},
		UserPoolId: aws.String(setting.GetUserPoolID()),
	})
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}
