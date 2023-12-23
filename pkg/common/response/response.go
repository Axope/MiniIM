package response

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}

func Fail(errMsg string) Response {
	return Response{
		Code: -1,
		Msg:  errMsg,
		Data: nil,
	}
}
