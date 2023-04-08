package handlers

import "github.com/sunxiaoning/restful-provider/mock/service"

type Handlers struct {
	service *service.Service
}

func NewHandlers(service *service.Service) *Handlers {
	return &Handlers{service: service}
}
