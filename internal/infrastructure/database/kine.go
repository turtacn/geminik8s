package database

import (
	"context"
	"os/exec"
	"syscall"

	"github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/internal/pkg/logger"
)

// KineManager manages the lifecycle of a Kine process.
type KineManager struct {
	log logger.Logger
	cmd *exec.Cmd
}

// NewKineManager creates a new Kine process manager.
func NewKineManager(log logger.Logger) *KineManager {
	return &KineManager{
		log: log,
	}
}

// Start launches the Kine process in the background.
// kinePath is the path to the kine binary.
// endpoint is the datastore endpoint (e.g., postgres://user:pass@host/db).
// listenAddress is the address for kine to listen on (e.g., unix:///var/run/kine.sock).
func (m *KineManager) Start(ctx context.Context, kinePath, endpoint, listenAddress string) error {
	if m.IsRunning() {
		m.log.Infof("Kine process is already running.")
		return nil
	}

	m.log.Infof("Starting Kine process...")
	m.log.Infof("  -> Kine Path: %s", kinePath)
	m.log.Infof("  -> Datastore Endpoint: %s", endpoint)
	m.log.Infof("  -> Listen Address: %s", listenAddress)

	m.cmd = exec.CommandContext(ctx, kinePath, "--endpoint", endpoint, "--listen-address", listenAddress)
	// Run in background
	m.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// In a real app, you'd handle stdout/stderr logging.
	// m.cmd.Stdout = ...
	// m.cmd.Stderr = ...

	if err := m.cmd.Start(); err != nil {
		return errors.Wrap(err, errors.OrchestratorError, "failed to start Kine process")
	}

	m.log.Infof("Kine process started with PID: %d", m.cmd.Process.Pid)
	return nil
}

// Stop terminates the Kine process.
func (m *KineManager) Stop() error {
	if !m.IsRunning() {
		m.log.Infof("Kine process is not running.")
		return nil
	}

	m.log.Infof("Stopping Kine process with PID: %d", m.cmd.Process.Pid)
	// Kill the entire process group to prevent orphaned processes
	if err := syscall.Kill(-m.cmd.Process.Pid, syscall.SIGKILL); err != nil {
		return errors.Wrap(err, errors.OrchestratorError, "failed to stop Kine process")
	}

	m.cmd.Wait() // Clean up zombie process
	m.cmd = nil
	m.log.Infof("Kine process stopped.")
	return nil
}

// IsRunning checks if the Kine process is currently running.
func (m *KineManager) IsRunning() bool {
	return m.cmd != nil && m.cmd.Process != nil && m.cmd.ProcessState == nil
}

//Personal.AI order the ending
