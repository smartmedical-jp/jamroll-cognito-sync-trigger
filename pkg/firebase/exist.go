package firebase

import (
	"context"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("firebase: user already exists")
)

func ExistByEmail(ctx context.Context, email string) (bool, error) {
	client, err := NewClient(ctx)
	if err != nil {
		return false, err
	}

	user, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}

	return true, nil
}
