package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetDemo(ctx echo.Context) error {
	return WriteResp(ctx, h.service.GetDemo())
}

func (h *Handlers) GetDemo400(ctx echo.Context) error {
	return WriteResp(ctx, h.service.GetDemo400())
}

func (h *Handlers) GetDemo401(ctx echo.Context) error {
	return WriteResp(ctx, h.service.GetDemo401())
}

func (h *Handlers) GetDemo403(ctx echo.Context) error {
	return WriteResp(ctx, h.service.GetDemo403())
}

func (h *Handlers) GetDemo404(ctx echo.Context) error {
	return WriteResp(ctx, h.service.GetDemo404())
}
