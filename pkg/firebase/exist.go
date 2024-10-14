package firebase

import (
	"context"
	"errors"
)

var (
	ErrUserNotExist     = errors.New("firebase: user not exists")
	ErrUserAlreadyExist = errors.New("firebase: user already exists")
)

func ExistByEmail(ctx context.Context, email string) (bool, error) {
	client, err := NewClient(ctx)
	if err != nil {
		return false, err
	}

	// Note: ユーザが存在しない場合にも err になる
	user, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, ErrUserNotExist
	}

	return true, nil
}
