package mail

import "bamboo-service/internal/service"

type sMail struct {
}

func init() {
	service.RegisterMail(New())
}

func New() *sMail {
	return &sMail{}
}
