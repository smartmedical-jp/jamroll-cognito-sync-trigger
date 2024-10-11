package ssm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"jam-roll-cognito-sync-trigger/pkg/aws/setting"
)

func NewClient() (*ssm.SSM, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(setting.GetRegion()),
	})
	if err != nil {
		return nil, err
	}

	return ssm.New(sess), nil
}
