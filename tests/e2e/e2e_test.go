//go:build e2e
// +build e2e

package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const binaryName = "gemin_k8s_e2e"

// TestMain is the entry point for e2e tests. It builds the binary once.
func TestMain(m *testing.M) {
	// Build the binary
	err := os.Chdir("../..")
	if err != nil {
		panic("failed to change dir to project root")
	}

	buildCmd := exec.Command("go", "build", "-o", filepath.Join("tests/e2e", binaryName), "./cmd/gemin_k8s")
	output, err := buildCmd.CombinedOutput()
	if err != nil {
		panic("failed to build binary for e2e tests: " + string(output))
	}

	// Run the tests
	code := m.Run()

	// Cleanup
	os.Remove(filepath.Join("tests/e2e", binaryName))
	os.Remove(filepath.Join("tests/e2e", "cluster.yaml"))

	os.Exit(code)
}

func TestInitCommand(t *testing.T) {
	// Change to the e2e test directory to run the command
	err := os.Chdir("tests/e2e")
	if err != nil {
		t.Fatalf("failed to change dir to e2e: %v", err)
	}
	defer os.Chdir("../..") // Go back to root at the end

	binaryPath := "./" + binaryName
	configPath := "./cluster.yaml"

	// Run the 'init' command
	cmd := exec.Command(binaryPath, "init",
		"--name=e2e-cluster",
		"--node1-ip=192.168.1.100",
		"--node2-ip=192.168.1.101",
		"--vip=192.168.1.200",
		"--config", configPath, // Override default config path for test isolation
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("init command failed: %v\nOutput: %s", err, string(output))
	}

	// Verify the output file was created
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatalf("config file '%s' was not created", configPath)
	}

	// Verify the content of the file
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("failed to read created config file: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "name: e2e-cluster") {
		t.Errorf("config file does not contain the correct cluster name")
	}
	if !strings.Contains(contentStr, "vip: 192.168.1.200") {
		t.Errorf("config file does not contain the correct VIP")
	}
}

//Personal.AI order the ending
