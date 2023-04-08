package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) HealthCheck(ctx echo.Context) error {
	ctx.NoContent(http.StatusNoContent)
	return nil
}
