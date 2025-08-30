package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger("debug", &buf, "text")

	if log == nil {
		t.Fatal("NewLogger returned nil")
	}

	log.Infof("hello %s", "world")

	output := buf.String()
	if !strings.Contains(output, "level=info") {
		t.Errorf("log output should contain 'level=info', got: %s", output)
	}
	if !strings.Contains(output, "msg=\"hello world\"") {
		t.Errorf("log output should contain 'msg=\"hello world\"', got: %s", output)
	}
}

func TestLogLevels(t *testing.T) {
	var buf bytes.Buffer

	t.Run("debug level", func(t *testing.T) {
		buf.Reset()
		log := NewLogger("debug", &buf, "text")
		log.Debugf("this is a debug message")
		if !strings.Contains(buf.String(), "level=debug") {
			t.Errorf("expected debug message to be logged at debug level")
		}
	})

	t.Run("info level", func(t *testing.T) {
		buf.Reset()
		log := NewLogger("info", &buf, "text")
		log.Debugf("this should not be logged")
		if buf.String() != "" {
			t.Errorf("expected no output for debug message at info level, got: %s", buf.String())
		}
	})
}

func TestJSONFormatter(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger("info", &buf, "json")

	log.Warnf("this is a warning")

	output := buf.String()
	// Basic check to see if it looks like JSON
	if !strings.HasPrefix(output, "{") || !strings.HasSuffix(output, "}\n") {
		t.Errorf("expected JSON output, got: %s", output)
	}
	if !strings.Contains(output, "\"level\":\"warning\"") {
		t.Errorf("expected JSON to contain level, got: %s", output)
	}
	if !strings.Contains(output, "\"msg\":\"this is a warning\"") {
		t.Errorf("expected JSON to contain message, got: %s", output)
	}
}

func TestWithField(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger("info", &buf, "text")

	log.WithField("request_id", "12345").Infof("user logged in")

	output := buf.String()
	if !strings.Contains(output, "request_id=12345") {
		t.Errorf("expected log to contain the field 'request_id', got: %s", output)
	}
}

func TestWithFields(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger("info", &buf, "text")

	fields := map[string]interface{}{
		"user_id":    "user-abc",
		"ip_address": "192.168.1.1",
	}
	log.WithFields(fields).Errorf("failed to process payment")

	output := buf.String()
	if !strings.Contains(output, "user_id=user-abc") {
		t.Errorf("expected log to contain the field 'user_id', got: %s", output)
	}
	if !strings.Contains(output, "ip_address=192.168.1.1") {
		t.Errorf("expected log to contain the field 'ip_address', got: %s", output)
	}
}

//Personal.AI order the ending
