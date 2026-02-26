package errors

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrUnauthorized       = New(401, "Unauthorized")
	ErrBadRequest         = New(400, "Bad Request")
	ErrNotFound           = New(404, "Not Found")
	ErrInternalServer     = New(500, "Internal Server Error")
	ErrInvalidCredentials = New(401, "用户名或密码错误")
	ErrTokenExpired       = New(401, "Refresh Token 已失效，请重新登录")
)
