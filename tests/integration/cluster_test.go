//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/turtacn/geminik8s/internal/app/config"
	"github.com/turtacn/geminik8s/internal/app/orchestrator"
	"github.com/turtacn/geminik8s/pkg/types"
	"github.com/turtacn/geminik8s/plugins/deploy"
)

// TestFullDeploymentWorkflow simulates a full deployment workflow,
// testing the integration between the orchestrator, plugins, and managers.
func TestFullDeploymentWorkflow(t *testing.T) {
	// 1. Setup
	configManager := config.NewManager()
	pluginManager := orchestrator.NewPluginManager()

	// For this test, the deploy plugin has no real dependencies.
	// In a more complex test, you would inject mocked infra clients here.
	deployPlugin := deploy.New()
	err := pluginManager.Register(deployPlugin)
	if err != nil {
		t.Fatalf("Failed to register deploy plugin: %v", err)
	}

	// The orchestrator doesn't use domain services for deploy in this implementation,
	// so we can pass nil for them.
	engine := orchestrator.NewEngine(pluginManager, configManager, nil)

	// 2. Create a test cluster config
	cfg := &types.ClusterConfig{
		APIVersion: "geminik8s.turtacn.com/v1alpha1",
		Kind:       "ClusterConfig",
		Metadata:   types.Metadata{Name: "integration-test-cluster"},
		Spec: types.ClusterSpec{
			Network: types.NetworkConfig{VIP: "10.0.0.1"},
			Nodes: []types.NodeInfo{
				{IP: "192.168.1.10", Role: types.RoleLeader},
				{IP: "192.168.1.11", Role: types.RoleFollower},
			},
		},
	}

	// 3. Execute
	t.Log("Starting deployment workflow...")
	err = engine.Deploy(context.Background(), cfg)

	// 4. Assert
	if err != nil {
		t.Fatalf("Deployment workflow failed: %v", err)
	}
	t.Log("Deployment workflow completed successfully.")
}

//Personal.AI order the ending
