package cognito

import (
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

type ExternalProvider string

const (
	Google    ExternalProvider = "Google"
	Microsoft ExternalProvider = "microsoft"
	Unknown   ExternalProvider = ""
)

func (e ExternalProvider) String() string {
	return string(e)
}

func GetExternalProvider(event events.CognitoEventUserPoolsPreSignup) ExternalProvider {
	if strings.Contains(event.UserName, "google") {
		return Google
	}
	if strings.Contains(event.UserName, "microsoft") {
		return Microsoft
	}
	return Unknown
}
