package firebase

import (
	"context"
	"errors"
)

var (
	ErrNotExistUser = errors.New("user does not exist")
)

func ExistByEmail(ctx context.Context, email string) error {
	client, err := NewClient(ctx)
	if err != nil {
		return err
	}

	user, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrNotExistUser
	}

	return validateUID(user.UID)
}
