package cognito

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"jam-roll-cognito-sync-trigger/pkg/aws/setting"
)

// NewIdpClient CognitoIdentityProvider クライアント生成
// - CognitoIdentityProvider は、AWS Cognito User Pools に関連する機能を提供します
func NewIdpClient() (*cognitoidentityprovider.CognitoIdentityProvider, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(setting.GetRegion()),
	})
	if err != nil {
		return nil, err
	}

	return cognitoidentityprovider.New(sess), nil
}
