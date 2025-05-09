package auth

import "bamboo-service/internal/service"

type sAuth struct {
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}
