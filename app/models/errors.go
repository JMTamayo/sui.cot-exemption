package models

// Error represents a generic error. It can be used across the application.
type Error struct {
	Details string `json:"details"`
}

// NewError creates a new error.
//
// Arguments:
//   - details: The details of the error.
//
// Returns:
//   - The error.
func NewError(details string) *Error {
	return &Error{
		Details: details,
	}
}

// Error returns the details of the error.
//
// Arguments:
//   - None.
//
// Returns:
//   - The details of the error.
func (e *Error) Error() string {
	return e.Details
}

// HTTPError represents an HTTP error. It can be used across the application.
type HTTPError struct {
	StatusCode int    `json:"-"`
	Error      *Error `json:"error"`
}

// NewHTTPError creates a new HTTP error.
//
// Arguments:
//   - statusCode: The status code of the error.
//   - error: The error.
//
// Returns:
//   - The HTTP error.
func NewHTTPError(statusCode int, body string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Error:      NewError(body),
	}
}
