package app

type Error struct {
	Error   error
	Message string
	Code    int
}

func NewError(msg string) *Error {
	return &Error{
		Error:   nil,
		Message: msg,
		Code:    0,
	}
}
