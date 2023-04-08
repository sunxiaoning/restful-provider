package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sunxiaoning/restful-provider/mock/config"
	"github.com/sunxiaoning/restful-provider/mock/handlers"
	"github.com/sunxiaoning/restful-provider/mock/service"
	"net/http"
)

const (
	HealthyCheckUrl = "/healthy"
)

func Init(config *config.Config, svr *service.Service) (*echo.Echo, error) {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		ctx.Logger().Errorf("handle http request: %v", err)
		err = handlers.WriteDefaultResp(ctx, code)
		if err != nil {
			ctx.Logger().Errorf("write default http resp: %v", err)
		}
	}
	h := handlers.NewHandlers(svr)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Skipper: func(c echo.Context) bool {
	//		if c.Request().URL.Path == HealthyCheckUrl {
	//			return true
	//		}
	//		return false
	//	},
	//}))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/demo", h.GetDemo)
	apiV1.GET("/demo400", h.GetDemo400)
	apiV1.GET("/demo401", h.GetDemo401)
	apiV1.GET("/demo403", h.GetDemo403)
	apiV1.GET("/demo404", h.GetDemo404)

	e.Any(HealthyCheckUrl, h.HealthCheck)
	return e, nil
}
