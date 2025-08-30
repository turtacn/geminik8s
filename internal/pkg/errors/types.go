package errors

import "fmt"

// ErrorCode defines the type for error codes.
type ErrorCode string

// Predefined error codes
const (
	// Unknown represents an unknown error.
	Unknown ErrorCode = "Unknown"

	// ConfigError represents an error related to configuration.
	ConfigError ErrorCode = "ConfigError"
	// NetworkError represents an error related to network operations.
	NetworkError ErrorCode = "NetworkError"
	// DatabaseError represents an error related to database operations.
	DatabaseError ErrorCode = "DatabaseError"
	// KubernetesError represents an error related to Kubernetes operations.
	KubernetesError ErrorCode = "KubernetesError"
	// PluginError represents an error related to a plugin.
	PluginError ErrorCode = "PluginError"
	// OrchestratorError represents an error within the orchestrator.
	OrchestratorError ErrorCode = "OrchestratorError"
	// ValidationError represents a validation error.
	ValidationError ErrorCode = "ValidationError"
	// IO Error represents a file system I/O error.
	IOError ErrorCode = "IOError"
)

// Error is a custom error type that includes a code, a message, and an optional underlying error.
type Error struct {
	Code    ErrorCode
	Message string
	Err     error
}

// Error returns the string representation of the error.
func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s - %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error for error chaining.
func (e *Error) Unwrap() error {
	return e.Err
}

// New creates a new error with a given code and message.
func New(code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Newf creates a new error with a given code and a formatted message.
func Newf(code ErrorCode, format string, a ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}

// Wrap wraps an existing error with a new error, adding a code and a message.
func Wrap(err error, code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Wrapf wraps an existing error with a new error, adding a code and a formatted message.
func Wrapf(err error, code ErrorCode, format string, a ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
		Err:     err,
	}
}

//Personal.AI order the ending
