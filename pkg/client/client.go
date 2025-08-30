package client

import (
	"context"

	"github.com/turtacn/geminik8s/internal/app/config"
	"github.com/turtacn/geminik8s/internal/app/orchestrator"
	"github.com/turtacn/geminik8s/internal/domain/cluster"
	"github.com/turtacn/geminik8s/internal/domain/node"
	"github.com/turtacn/geminik8s/internal/domain/storage"
	"github.com/turtacn/geminik8s/pkg/api"
	"github.com/turtacn/geminik8s/pkg/types"
)

// Client is the interface for the geminik8s programmatic client.
type Client interface {
	Deploy(ctx context.Context, cfg *types.ClusterConfig) error
	GetStatus(ctx context.Context, cfg *types.ClusterConfig) (*types.ClusterStatus, error)
	// Add other high-level methods here as needed.
}

// client implements the Client interface.
type client struct {
	orchestrator api.Orchestrator
}

// Config holds the configuration for creating a new client.
type Config struct {
	// KubeconfigPath could be a config option for the client.
	// KubeconfigPath string
}

// NewClient creates a new geminik8s client.
// This function is responsible for wiring up all the application dependencies.
func NewClient(clientCfg *Config) (Client, error) {
	// 1. Initialize Managers
	configManager := config.NewManager()
	pluginManager := orchestrator.NewPluginManager()
	// TODO: Register real plugins
	// pluginManager.Register(deploy.New(...))
	// pluginManager.Register(health.New(...))

	// 2. Initialize Infrastructure (this is the tricky part for a pure SDK)
	// In a real scenario, the client config would need DB connection info, etc.
	// For now, we use nil dependencies for the domain services.
	// dbClient := database.NewPostgresClient(...)
	// k8sClient, _ := kubernetes.NewK8sClient(clientCfg.KubeconfigPath)

	// 3. Initialize Domain Services
	// These would take real infrastructure clients.
	nodeSvc := node.NewService(nil, nil, nil)
	storageSvc := storage.NewService(nil, nil)
	clusterSvc := cluster.NewService(nil, nodeSvc, storageSvc)

	// 4. Initialize Orchestrator
	orch := orchestrator.NewEngine(pluginManager, configManager, clusterSvc)

	return &client{
		orchestrator: orch,
	}, nil
}

// Deploy deploys a cluster.
func (c *client) Deploy(ctx context.Context, cfg *types.ClusterConfig) error {
	return c.orchestrator.Deploy(ctx, cfg)
}

// GetStatus retrieves the status of a cluster.
func (c *client) GetStatus(ctx context.Context, cfg *types.ClusterConfig) (*types.ClusterStatus, error) {
	return c.orchestrator.GetStatus(ctx, cfg)
}

//Personal.AI order the ending
