package registry

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/pkg/firebase"
)

type Registry struct {
	firebase *auth.Client
}

func NewRegistry(ctx context.Context) (*Registry, error) {
	fb, err := firebase.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	rgst := &Registry{
		firebase: fb,
	}

	return rgst, nil
}
