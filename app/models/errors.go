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
