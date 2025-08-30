package system

import (
	"os"
	"os/exec"

	"github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/api"
)

// systemOperator implements the api.SystemOperator interface.
type systemOperator struct{}

// NewSystemOperator creates a new system operator.
func NewSystemOperator() api.SystemOperator {
	return &systemOperator{}
}

// RunCommand executes a shell command and returns its combined output.
func (o *systemOperator) RunCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), errors.Wrapf(err, errors.OrchestratorError, "command failed: %s %v", command, args)
	}
	return string(output), nil
}

// WriteFile writes data to a file.
func (o *systemOperator) WriteFile(path string, content []byte, perm os.FileMode) error {
	err := os.WriteFile(path, content, perm)
	if err != nil {
		return errors.Wrapf(err, errors.IOError, "failed to write file: %s", path)
	}
	return nil
}

// ReadFile reads data from a file.
func (o *systemOperator) ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, errors.IOError, "failed to read file: %s", path)
	}
	return data, nil
}

//Personal.AI order the ending
