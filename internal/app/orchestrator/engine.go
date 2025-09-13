package orchestrator

import (
	"context"
	"errors"

	"github.com/turtacn/geminik8s/internal/domain/cluster"
	"github.com/turtacn/geminik8s/pkg/api"
	"github.com/turtacn/geminik8s/pkg/types"
)

// engine implements the api.Orchestrator interface.
type engine struct {
	pluginManager api.PluginManager
	configManager api.ConfigManager
	clusterSvc    *cluster.Service
	// other domain services would be injected here
}

// NewEngine creates a new orchestrator engine.
func NewEngine(
	pluginMgr api.PluginManager,
	configMgr api.ConfigManager,
	clusterSvc *cluster.Service,
) api.Orchestrator {
	return &engine{
		pluginManager: pluginMgr,
		configManager: configMgr,
		clusterSvc:    clusterSvc,
	}
}

// Init is a bit of a special case, as it doesn't operate on an existing cluster,
// but creates the initial configuration.
func (e *engine) Init(ctx context.Context, cfg *types.ClusterConfig) error {
	// The logic for init is mostly about generating the config file.
	// We can use the config manager to save the generated config.
	// The user of the orchestrator (e.g., the CLI command) would be responsible
	// for building the initial ClusterConfig object from flags/prompts.
	return e.configManager.Save(cfg, "cluster.yaml") //
}

// Deploy orchestrates the deployment of a cluster using a plugin.
func (e *engine) Deploy(ctx context.Context, cfg *types.ClusterConfig) error {
	// First, create the cluster record in the database using the domain service.
	if e.clusterSvc != nil {
		_, err := e.clusterSvc.CreateCluster(ctx, cfg)
		if err != nil {
			return err
		}
	}

	// Then, execute the deployment plugin.
	params := api.PluginParams{
		"config": cfg,
	}
	_, err := e.pluginManager.Execute(ctx, "deploy", params)
	return err
}

// GetStatus gets the status of a cluster.
func (e *engine) GetStatus(ctx context.Context, cfg *types.ClusterConfig) (*types.ClusterStatus, error) {
	// This could also be a plugin.
	params := api.PluginParams{
		"config": cfg,
	}
	result, err := e.pluginManager.Execute(ctx, "health", params)
	if err != nil {
		return nil, err
	}

	statusStr, ok := result.Data["status"].(string)
	if !ok {
		s := types.StatusUnknown
		return &s, nil
	}
	s := types.ClusterStatus(statusStr)
	return &s, nil
}

// The rest of the methods would follow a similar pattern,
// typically finding the right plugin and executing it with the given config.

func (e *engine) Failover(ctx context.Context, cfg *types.ClusterConfig, promoteNode string) error {
	return errors.New("not implemented")
}

func (e *engine) Upgrade(ctx context.Context, cfg *types.ClusterConfig, version string) error {
	return errors.New("not implemented")
}

func (e *engine) ReplaceNode(ctx context.Context, cfg *types.ClusterConfig, oldNode, newNode string) error {
	return errors.New("not implemented")
}

func (e *engine) Backup(ctx context.Context, cfg *types.ClusterConfig, destination string) error {
	return errors.New("not implemented")
}

func (e *engine) Restore(ctx context.Context, cfg *types.ClusterConfig, source string) error {
	return errors.New("not implemented")
}

//Personal.AI order the ending
