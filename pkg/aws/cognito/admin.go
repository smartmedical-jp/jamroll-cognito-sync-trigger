package cognito

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"jam-roll-cognito-sync-trigger/pkg/aws/setting"
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
	_, err = client.AdminLinkProviderForUser(&cognitoidentityprovider.AdminLinkProviderForUserInput{
		DestinationUser: &cognitoidentityprovider.ProviderUserIdentifierType{
			ProviderAttributeValue: aws.String(*destinationUser.Username),
			ProviderName:           aws.String("Cognito"),
		},
		SourceUser: &cognitoidentityprovider.ProviderUserIdentifierType{
			ProviderAttributeName: aws.String("Cognito_Subject"),
			// TODO 上書きされるのか確認（多分されそう）
			ProviderAttributeValue: aws.String(event.UserName),
			ProviderName:           aws.String(provider.String()),
		},
		UserPoolId: aws.String(setting.GetUserPoolID()),
	})
	if err != nil {
		return err
	}

	return nil
}
