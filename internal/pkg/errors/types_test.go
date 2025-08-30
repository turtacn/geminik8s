package errors

import (
	"errors"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(ValidationError, "validation failed")
	if err.Code != ValidationError {
		t.Errorf("expected code %s, got %s", ValidationError, err.Code)
	}
	if err.Message != "validation failed" {
		t.Errorf("expected message 'validation failed', got '%s'", err.Message)
	}
	if err.Err != nil {
		t.Errorf("expected underlying error to be nil, got %v", err.Err)
	}
}

func TestNewf(t *testing.T) {
	err := Newf(ConfigError, "config file not found at %s", "/path/to/config")
	expectedMsg := "config file not found at /path/to/config"
	if err.Message != expectedMsg {
		t.Errorf("expected message '%s', got '%s'", expectedMsg, err.Message)
	}
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, DatabaseError, "failed to query users")

	if wrappedErr.Code != DatabaseError {
		t.Errorf("expected code %s, got %s", DatabaseError, wrappedErr.Code)
	}
	if wrappedErr.Message != "failed to query users" {
		t.Errorf("expected message 'failed to query users', got '%s'", wrappedErr.Message)
	}
	if wrappedErr.Err != originalErr {
		t.Errorf("expected underlying error to be the original error")
	}
}

func TestErrorMethod(t *testing.T) {
	t.Run("without underlying error", func(t *testing.T) {
		err := New(NetworkError, "timeout")
		expected := "NetworkError: timeout"
		if err.Error() != expected {
			t.Errorf("expected error string '%s', got '%s'", expected, err.Error())
		}
	})

	t.Run("with underlying error", func(t *testing.T) {
		originalErr := errors.New("connection refused")
		err := Wrap(originalErr, NetworkError, "failed to connect")
		expectedPrefix := "NetworkError: failed to connect"
		expectedSuffix := "connection refused"

		if !strings.HasPrefix(err.Error(), expectedPrefix) {
			t.Errorf("error string should have prefix '%s', got '%s'", expectedPrefix, err.Error())
		}
		if !strings.HasSuffix(err.Error(), expectedSuffix) {
			t.Errorf("error string should have suffix '%s', got '%s'", expectedSuffix, err.Error())
		}
	})
}

func TestUnwrap(t *testing.T) {
	originalErr := errors.New("root cause")
	wrappedErr := Wrap(originalErr, Unknown, "an error occurred")

	if unwrapped := errors.Unwrap(wrappedErr); unwrapped != originalErr {
		t.Errorf("errors.Unwrap did not return the original error")
	}
}

//Personal.AI order the ending
