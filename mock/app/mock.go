package app

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sunxiaoning/restful-provider/mock/config"
	"github.com/sunxiaoning/restful-provider/mock/router"
	"github.com/sunxiaoning/restful-provider/mock/service"
	"net/http"
)

type Mock struct {
	config *config.Config
	server *http.Server
}

func NewSnapshot(cfg *config.Config) (*Mock, error) {
	s := &Mock{config: cfg}
	svr := service.NewService(cfg)
	rou, err := router.Init(cfg, svr)
	if err != nil {
		return nil, err
	}
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: rou,
	}

	return s, nil
}

func (s *Mock) Serve() error {
	logrus.Infof("started rest server at %d", s.config.Server.Port)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logrus.Info("rest server exited!")
			return nil
		}
		logrus.Fatalf("rest server closed unexpect: %v", err)
	}
	return nil
}

func (s *Mock) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		logrus.WithField("err", err).Errorf("rest server failed to stop!")
	}
	logrus.Info("rest server closed under request")
}
