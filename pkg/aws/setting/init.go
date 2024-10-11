package setting

import "github.com/aws/aws-lambda-go/events"

type cognitoEvent interface {
	getRegion() string
	getUserPoolID() string
}

func InitSetting[T cognitoEvent](event T) {
	setRegion(event.getRegion())
	setUserPoolID(event.getUserPoolID())
}

type (
	MigrateUser struct {
		Event events.CognitoEventUserPoolsMigrateUser
	}
	PreSignup struct {
		Event events.CognitoEventUserPoolsPreSignup
	}
)

func (e MigrateUser) getRegion() string {
	return e.Event.Region
}

func (e MigrateUser) getUserPoolID() string {
	return e.Event.UserPoolID
}

func (e PreSignup) getRegion() string {
	return e.Event.Region
}

func (e PreSignup) getUserPoolID() string {
	return e.Event.UserPoolID
}
