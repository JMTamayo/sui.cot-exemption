package models

type Error struct {
	Details string `json:"details"`
}

func NewError(details string) *Error {
	return &Error{
		Details: details,
	}
}

func (e *Error) Error() string {
	return e.Details
}

type HttpError struct {
	StatusCode int
	Body       string
}

func NewHttpError(statusCode int, body string) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Body:       body,
	}
}
