package role

import "bamboo-service/internal/service"

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}
