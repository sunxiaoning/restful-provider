package service

import (
	"encoding/json"
	"strconv"
)

type BizErr struct {
	Code int
	Msg  string
}

func NewBizErr(httpCode, bizCode int, msg string) *BizErr {
	code, _ := strconv.Atoi(strconv.Itoa(httpCode) + strconv.Itoa(bizCode))
	return &BizErr{Code: code, Msg: msg}
}

func (b BizErr) Error() string {
	be, _ := json.Marshal(b)
	return string(be)
}
