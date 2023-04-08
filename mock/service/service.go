package service

import "github.com/sunxiaoning/restful-provider/mock/config"

type Service struct {
	config *config.Config
}

func NewService(config *config.Config) *Service {
	return &Service{config: config}
}
