package logerror

import "fmt"

// LogError is a custom error type that includes an error message, the underlying error, and an error type.
type LogError struct {
	Err       error
	Msg       string
	ErrorType string
}

// Error implements the error interface for LogError.
func (e LogError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorType, e.Msg)
}

// Unwrap provides compatibility for errors.Is and errors.As
func (e LogError) Unwrap() error {
	return e.Err
}

// NewBusinessError creates a new LogError with the "business" error type.
// If the innerErr is already a LogError, it reuses it instead of nesting.
func NewBusinessError(message string, innerErr error) error {
	if _, ok := innerErr.(LogError); ok {
		return innerErr
	}
	return LogError{
		Err:       innerErr,
		Msg:       message,
		ErrorType: "business",
	}
}

// NewDatabaseError creates a new LogError with the "database" error type.
// If the innerErr is already a LogError, it reuses it instead of nesting.
func NewDatabaseError(message string, innerErr error) error {
	if _, ok := innerErr.(LogError); ok {
		return innerErr
	}
	return LogError{
		Err:       innerErr,
		Msg:       message,
		ErrorType: "database",
	}
}
