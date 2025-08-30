package orchestrator

import (
	"context"
	"sync"

	"github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/api"
)

// pluginManager implements the api.PluginManager interface.
type pluginManager struct {
	plugins map[string]api.Plugin
	mu      sync.RWMutex
}

// NewPluginManager creates a new plugin manager.
func NewPluginManager() api.PluginManager {
	return &pluginManager{
		plugins: make(map[string]api.Plugin),
	}
}

// Register adds a new plugin to the manager.
func (pm *pluginManager) Register(plugin api.Plugin) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, exists := pm.plugins[plugin.Name()]; exists {
		return errors.Newf(errors.PluginError, "plugin with name '%s' already registered", plugin.Name())
	}
	pm.plugins[plugin.Name()] = plugin
	// In a real app, you would log this.
	// log.Infof("Registered plugin: %s (version: %s)", plugin.Name(), plugin.Version())
	return nil
}

// Get retrieves a plugin by name.
func (pm *pluginManager) Get(name string) (api.Plugin, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	plugin, exists := pm.plugins[name]
	if !exists {
		return nil, errors.Newf(errors.PluginError, "plugin with name '%s' not found", name)
	}
	return plugin, nil
}

// Execute finds a plugin by name and executes it.
func (pm *pluginManager) Execute(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
	plugin, err := pm.Get(name)
	if err != nil {
		return nil, err
	}

	if err := plugin.Validate(params); err != nil {
		return nil, errors.Wrapf(err, errors.ValidationError, "plugin '%s' validation failed", name)
	}

	result, err := plugin.Execute(ctx, params)
	if err != nil {
		return nil, errors.Wrapf(err, errors.PluginError, "plugin '%s' execution failed", name)
	}

	return result, nil
}

//Personal.AI order the ending
