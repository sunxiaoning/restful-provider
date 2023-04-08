package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sunxiaoning/restful-provider/mock/service"
	"net/http"
	"strconv"
)

type RestResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var defaultErrResp = &RestResponse{
	Code: 5000000,
	Msg:  "internal server error",
}

func NewRespWithCode(code int, msg string, data interface{}) *RestResponse {
	return &RestResponse{Code: code, Msg: msg, Data: data}
}

func NewRestResponse(data interface{}) *RestResponse {
	return &RestResponse{Data: data}
}

func WriteDefaultResp(ctx echo.Context, httpCode int) error {
	if httpCode == http.StatusInternalServerError {
		return ctx.JSON(http.StatusInternalServerError, defaultErrResp)
	}
	code, err := strconv.Atoi(strconv.Itoa(httpCode) + "0000")
	if err != nil {
		ctx.Logger().Errorf("set biz code:%d, %w", httpCode, err)
	}
	msg := http.StatusText(httpCode)
	if msg == "" {
		msg = "unknown error!"
	}
	return ctx.JSON(httpCode, NewRespWithCode(code, msg, nil))
}

func WriteResp(ctx echo.Context, resp service.Resp) error {
	if resp.Error == nil {
		ctx.JSON(http.StatusOK, NewRestResponse(resp.Data))
		return nil
	}
	bizErr, ok := resp.Error.(*service.BizErr)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, defaultErrResp)
		return nil
	}
	httpCode, err := strconv.Atoi(strconv.Itoa(bizErr.Code)[0:3])
	if err != nil {
		ctx.Logger().Errorf("get biz code:%d, %w", httpCode, err)
	}
	return ctx.JSON(httpCode, NewRespWithCode(bizErr.Code, bizErr.Msg, resp.Data))
}
