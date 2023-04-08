package service

import (
	"net/http"
	"time"
)

func (s *Service) GetDemo() Resp {
	return newResp(&Demo{
		Id:        1,
		Name:      "demo01",
		Age:       10,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})
}

func (s *Service) GetDemo400() Resp {
	return newErrResp(NewBizErr(http.StatusBadRequest, 1000, "id param invalid!"))
}

func (s *Service) GetDemo401() Resp {
	return newErrResp(NewBizErr(http.StatusUnauthorized, 1000, "user need login!"))
}

func (s *Service) GetDemo403() Resp {
	return newErrResp(NewBizErr(http.StatusForbidden, 1000, "user need auth!"))
}

func (s *Service) GetDemo404() Resp {
	return newErrResp(NewBizErr(http.StatusNotFound, 1000, "demo record not found!"))
}

type Demo struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
