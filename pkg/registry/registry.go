package registry

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"jam-roll-cognito-sync-trigger/pkg/firebase"
)

// NOTE: 現状（2024/10/09）、以下コード使っていない

type Registry struct {
	Firebase *auth.Client
}

func NewRegistry(ctx context.Context) (*Registry, error) {
	fb, err := firebase.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	rgst := &Registry{
		Firebase: fb,
	}

	return rgst, nil
}
