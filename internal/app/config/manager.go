package config

import (
	"os"

	"github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/internal/pkg/utils"
	"github.com/turtacn/geminik8s/pkg/api"
	"github.com/turtacn/geminik8s/pkg/types"
	"sigs.k8s.io/yaml"
)

// Manager implements the api.ConfigManager interface.
type Manager struct {
	// In a real app, you'd inject a logger here.
}

// NewManager creates a new configuration manager.
func NewManager() api.ConfigManager {
	return &Manager{}
}

// Load reads a cluster configuration file from a given path.
func (m *Manager) Load(path string) (*types.ClusterConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, errors.IOError, "failed to read config file: %s", path)
	}

	var cfg types.ClusterConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.Wrapf(err, errors.ConfigError, "failed to parse config file: %s", path)
	}

	if err := m.Validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Save writes a cluster configuration to a given path.
func (m *Manager) Save(cfg *types.ClusterConfig, path string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, errors.ConfigError, "failed to marshal config")
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return errors.Wrapf(err, errors.IOError, "failed to write config file: %s", path)
	}

	return nil
}

// Validate checks if the cluster configuration is valid.
func (m *Manager) Validate(cfg *types.ClusterConfig) error {
	if cfg.APIVersion == "" || cfg.Kind == "" {
		return errors.New(errors.ValidationError, "apiVersion and kind must be set")
	}
	if cfg.Metadata.Name == "" {
		return errors.New(errors.ValidationError, "metadata.name must be set")
	}
	if len(cfg.Spec.Nodes) != 2 {
		return errors.New(errors.ValidationError, "exactly two nodes must be defined in spec.nodes")
	}
	if cfg.Spec.Network.VIP == "" {
		return errors.New(errors.ValidationError, "spec.network.vip must be set")
	}
	// Add more validation rules here...
	return nil
}

// Render renders a configuration template.
func (m *Manager) Render(templatePath string, data interface{}) (string, error) {
	templateContent, err := utils.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	return utils.RenderTemplate(templateContent, data)
}

//Personal.AI order the ending
