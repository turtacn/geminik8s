package api

import (
	"context"
	"os"

	"github.com/turtacn/geminik8s/pkg/types"
)

// Orchestrator defines the interface for the main engine that drives all operations.
type Orchestrator interface {
	Init(ctx context.Context, cfg *types.ClusterConfig) error
	Deploy(ctx context.Context, cfg *types.ClusterConfig) error
	GetStatus(ctx context.Context, cfg *types.ClusterConfig) (*types.ClusterStatus, error)
	Failover(ctx context.Context, cfg *types.ClusterConfig, promoteNode string) error
	Upgrade(ctx context.Context, cfg *types.ClusterConfig, version string) error
	ReplaceNode(ctx context.Context, cfg *types.ClusterConfig, oldNode, newNode string) error
	Backup(ctx context.Context, cfg *types.ClusterConfig, destination string) error
	Restore(ctx context.Context, cfg *types.ClusterConfig, source string) error
}

// PluginParams is a map for passing parameters to a plugin.
type PluginParams map[string]interface{}

// PluginResult holds the result of a plugin execution.
type PluginResult struct {
	Success bool
	Message string
	Data    map[string]interface{}
}

// Plugin defines the interface for all plugins in the system.
type Plugin interface {
	Name() string
	Version() string
	Execute(ctx context.Context, params PluginParams) (*PluginResult, error)
	Validate(params PluginParams) error
	Cleanup(ctx context.Context) error
}

// PluginManager defines the interface for managing plugins.
type PluginManager interface {
	Register(plugin Plugin) error
	Get(name string) (Plugin, error)
	Execute(ctx context.Context, name string, params PluginParams) (*PluginResult, error)
}

// ConfigManager defines the interface for managing cluster configurations.
type ConfigManager interface {
	Load(path string) (*types.ClusterConfig, error)
	Save(cfg *types.ClusterConfig, path string) error
	Validate(cfg *types.ClusterConfig) error
	Render(templatePath string, data interface{}) (string, error)
}

// Infrastructure interfaces

// DBClient defines the interface for database operations.
type DBClient interface {
	Connect() error
	Close() error
	Execute(query string, args ...interface{}) error
	Query(query string, args ...interface{}) (interface{}, error) // Simplified
}

// K8sClient defines the interface for interacting with the Kubernetes API.
type K8sClient interface {
	GetNodes(ctx context.Context) ([]types.Node, error)
	Apply(ctx context.Context, manifest []byte) error
	Delete(ctx context.Context, manifest []byte) error
}

// SystemOperator defines the interface for system-level operations.
type SystemOperator interface {
	RunCommand(command string, args ...string) (string, error)
	WriteFile(path string, content []byte, perm os.FileMode) error
	ReadFile(path string) ([]byte, error)
}

// NetworkOperator defines the interface for network-related operations.
type NetworkOperator interface {
	CheckConnectivity(host string, port int) error
	ManageVIP(action string, vip string) error // e.g., action="add" or "del"
}

//Personal.AI order the ending
