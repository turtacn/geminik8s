//go:build e2e
// +build e2e

package e2e

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const (
	binaryName      = "gemin_k8s_e2e"
	testNetworkName = "geminik8s-e2e-test-net"
)

// TestMain sets up and tears down the E2E test environment.
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

	// Setup Docker network
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic("failed to create docker client: " + err.Error())
	}
	_, err = cli.NetworkCreate(context.Background(), testNetworkName, types.NetworkCreate{})
	if err != nil {
		panic("failed to create docker network: " + err.Error())
	}

	// Run tests
	code := m.Run()

	// Cleanup
	cli.NetworkRemove(context.Background(), testNetworkName)
	os.Remove(filepath.Join("tests/e2e", binaryName))
	os.RemoveAll(filepath.Join("tests/e2e", "e2e-cluster")) // Clean up config dir

	os.Exit(code)
}

func TestInitCommand(t *testing.T) {
	// Change to the e2e test directory to run the command
	err := os.Chdir("tests/e2e")
	if err != nil {
		// If we are already in the correct directory, we can ignore the error.
		if !strings.HasSuffix(err.Error(), "no such file or directory") {
			t.Fatalf("failed to change dir to e2e: %v", err)
		}
	}
	defer os.Chdir("../..") // Go back to root at the end

	binaryPath := "./" + binaryName
	configDir := "./e2e-cluster"

	// Run the 'init' command
	cmd := exec.Command(binaryPath, "init",
		"--name=e2e-cluster",
		"--node1-ip=172.28.0.2",
		"--node2-ip=172.28.0.3",
		"--vip=172.28.0.100",
		"--config-dir", configDir,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("init command failed: %v\nOutput: %s", err, string(output))
	}

	// Verify the output file was created
	configPath := filepath.Join(configDir, "cluster.yaml")
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
	if !strings.Contains(contentStr, "vip: 172.28.0.100") {
		t.Errorf("config file does not contain the correct VIP")
	}
}

func TestDeployCommand(t *testing.T) {
	// This test is a work in progress and does not yet perform a full deployment.
	// A full implementation would require a more sophisticated test harness.
	t.Log("E2E deploy test is a work-in-progress.")
}

// Helper function to manage containers (a more complete version would be needed)
func startTestNode(t *testing.T, cli *client.Client, name, image string) string {
	ctx := context.Background()

	// Pull image
	reader, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		t.Fatalf("failed to pull image %s: %v", image, err)
	}
	io.Copy(os.Stdout, reader)

	// Create container
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		Tty:   true,
	}, &container.HostConfig{
		NetworkMode: container.NetworkMode(testNetworkName),
		Privileged:  true, // Required for K3s
	}, nil, nil, name)
	if err != nil {
		t.Fatalf("failed to create container %s: %v", name, err)
	}

	// Start container
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		t.Fatalf("failed to start container %s: %v", name, err)
	}

	// Get container IP
	inspect, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		t.Fatalf("failed to inspect container %s: %v", name, err)
	}
	ip := inspect.NetworkSettings.Networks[testNetworkName].IPAddress

	t.Logf("Started container %s with IP %s", name, ip)
	return ip
}
