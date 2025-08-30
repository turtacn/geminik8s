package utils

import (
	"bytes"
	"os"
	"text/template"
	"time"

	"github.com/turtacn/geminik8s/internal/pkg/errors"
)

// FileExists checks if a file exists at the given path.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadFile reads the content of a file into a string.
func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", errors.Wrapf(err, errors.IOError, "failed to read file: %s", path)
	}
	return string(data), nil
}

// RenderTemplate renders a template with the given data.
func RenderTemplate(tmplContent string, data interface{}) (string, error) {
	tmpl, err := template.New("config").Parse(tmplContent)
	if err != nil {
		return "", errors.Wrap(err, errors.ConfigError, "failed to parse template")
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", errors.Wrap(err, errors.ConfigError, "failed to render template")
	}

	return buf.String(), nil
}

// Retry executes a function multiple times until it succeeds or the max attempts are reached.
func Retry(attempts int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(sleep)
	}
	return errors.Wrapf(err, errors.Unknown, "failed after %d attempts", attempts)
}

//Personal.AI order the ending
