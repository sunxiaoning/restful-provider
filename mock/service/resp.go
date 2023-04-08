package service

type Resp struct {
	Data  interface{}
	Error error
}

func newResp(data interface{}) Resp {
	return Resp{Data: data}
}

func newErrResp(err error) Resp {
	return Resp{Error: err}
}
