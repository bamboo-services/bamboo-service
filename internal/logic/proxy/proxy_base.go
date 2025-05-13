package proxy

import "bamboo-service/internal/service"

type sProxy struct {
}

func init() {
	service.RegisterProxy(New())
}

func New() *sProxy {
	return &sProxy{}
}
