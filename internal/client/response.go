package client

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(status int, message string, data interface{}) *Response {
	return &Response{Status: int(status), Message: message, Data: data}
}
