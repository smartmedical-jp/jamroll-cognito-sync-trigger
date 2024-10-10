package firebase

import (
	"errors"
	"strings"
)

var (
	ErrEmptyUID = errors.New("uid is empty")
)

func validateUID(uid string) error {
	strings.TrimSpace(uid)
	if uid == "" {
		return ErrEmptyUID
	}

	return nil
}
