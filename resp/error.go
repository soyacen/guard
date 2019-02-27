package resp

func NewError(code int, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *Error) Error() string {
	return e.Message
}
