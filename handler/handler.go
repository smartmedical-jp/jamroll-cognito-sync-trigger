package handler

import "github.com/smartmedical-jp/jam-roll-cognito-sync-trigger/pkg/registry"

type Handler struct {
	registry registry.Registry
}

func NewHandler(registry registry.Registry) *Handler {
	return &Handler{
		registry: registry,
	}
}
