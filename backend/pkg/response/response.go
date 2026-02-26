package response

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func Error(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

func Unauthorized(message string) *Response {
	return Error(http.StatusUnauthorized, message)
}

func BadRequest(message string) *Response {
	return Error(http.StatusBadRequest, message)
}

func NotFound(message string) *Response {
	return Error(http.StatusNotFound, message)
}

func InternalServerError(message string) *Response {
	return Error(http.StatusInternalServerError, message)
}
