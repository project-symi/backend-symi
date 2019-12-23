package controllers

type Error struct {
	Message string
}

func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}

func ValidationError(info string, err error) *Error {
	return &Error{
		Message: "Validation error for " + info + ": " + err.Error(),
	}
}
