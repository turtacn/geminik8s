package orchestrator

import (
	"context"
	"testing"

	"github.com/turtacn/geminik8s/pkg/api"
	"github.com/turtacn/geminik8s/pkg/types"
)

// --- Mocks ---

type mockPluginManager struct {
	ExecuteFunc func(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error)
	// Add other methods if needed for other tests
}

func (m *mockPluginManager) Register(plugin api.Plugin) error    { return nil }
func (m *mockPluginManager) Get(name string) (api.Plugin, error) { return nil, nil }
func (m *mockPluginManager) Execute(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
	return m.ExecuteFunc(ctx, name, params)
}

type mockConfigManager struct {
	LoadFunc     func(path string) (*types.ClusterConfig, error)
	SaveFunc     func(cfg *types.ClusterConfig, path string) error
	ValidateFunc func(cfg *types.ClusterConfig) error
	RenderFunc   func(templatePath string, data interface{}) (string, error)
}

func (m *mockConfigManager) Load(path string) (*types.ClusterConfig, error) { return m.LoadFunc(path) }
func (m *mockConfigManager) Save(cfg *types.ClusterConfig, path string) error {
	return m.SaveFunc(cfg, path)
}
func (m *mockConfigManager) Validate(cfg *types.ClusterConfig) error { return m.ValidateFunc(cfg) }
func (m *mockConfigManager) Render(templatePath string, data interface{}) (string, error) {
	return m.RenderFunc(templatePath, data)
}

// --- Tests ---

func TestEngineDeploy(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		executeCalled := false
		pluginName := ""
		mockPluginMgr := &mockPluginManager{
			ExecuteFunc: func(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
				executeCalled = true
				pluginName = name
				if _, ok := params["config"]; !ok {
					t.Error("expected 'config' in plugin params")
				}
				return &api.PluginResult{Success: true}, nil
			},
		}

		engine := NewEngine(mockPluginMgr, nil, nil)
		cfg := &types.ClusterConfig{Metadata: types.Metadata{Name: "test"}}

		err := engine.Deploy(context.Background(), cfg)
		if err != nil {
			t.Fatalf("Deploy failed: %v", err)
		}

		if !executeCalled {
			t.Errorf("expected plugin manager Execute to be called")
		}
		if pluginName != "deploy" {
			t.Errorf("expected 'deploy' plugin to be called, got '%s'", pluginName)
		}
	})

	t.Run("PluginError", func(t *testing.T) {
		mockPluginMgr := &mockPluginManager{
			ExecuteFunc: func(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
				return nil, errors.New("plugin failed")
			},
		}

		engine := NewEngine(mockPluginMgr, nil, nil)
		cfg := &types.ClusterConfig{Metadata: types.Metadata{Name: "test"}}

		err := engine.Deploy(context.Background(), cfg)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() != "plugin failed" {
			t.Fatalf("expected 'plugin failed' error, got '%v'", err)
		}
	})
}

func TestEngineGetStatus(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockPluginMgr := &mockPluginManager{
			ExecuteFunc: func(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
				if name != "health" {
					t.Errorf("expected 'health' plugin to be called, got '%s'", name)
				}
				return &api.PluginResult{
					Success: true,
					Data:    map[string]interface{}{"status": "Running"},
				}, nil
			},
		}

		engine := NewEngine(mockPluginMgr, nil, nil)
		cfg := &types.ClusterConfig{Metadata: types.Metadata{Name: "test"}}

		status, err := engine.GetStatus(context.Background(), cfg)
		if err != nil {
			t.Fatalf("GetStatus failed: %v", err)
		}

		if *status != types.StatusRunning {
			t.Errorf("expected status to be 'Running', got '%s'", *status)
		}
	})

	t.Run("PluginError", func(t *testing.T) {
		mockPluginMgr := &mockPluginManager{
			ExecuteFunc: func(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
				return nil, errors.New("plugin failed")
			},
		}

		engine := NewEngine(mockPluginMgr, nil, nil)
		cfg := &types.ClusterConfig{Metadata: types.Metadata{Name: "test"}}

		_, err := engine.GetStatus(context.Background(), cfg)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() != "plugin failed" {
			t.Fatalf("expected 'plugin failed' error, got '%v'", err)
		}
	})

	t.Run("InvalidData", func(t *testing.T) {
		mockPluginMgr := &mockPluginManager{
			ExecuteFunc: func(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
				return &api.PluginResult{
					Success: true,
					Data:    map[string]interface{}{"status": 123}, // Invalid data type
				}, nil
			},
		}

		engine := NewEngine(mockPluginMgr, nil, nil)
		cfg := &types.ClusterConfig{Metadata: types.Metadata{Name: "test"}}

		status, err := engine.GetStatus(context.Background(), cfg)
		if err != nil {
			t.Fatalf("GetStatus failed with unexpected error: %v", err)
		}
		if *status != types.StatusUnknown {
			t.Errorf("expected status to be 'Unknown', got '%s'", *status)
		}
	})
}

func TestEngineInit(t *testing.T) {
	saveCalled := false
	var savedPath string
	mockConfigMgr := &mockConfigManager{
		SaveFunc: func(cfg *types.ClusterConfig, path string) error {
			saveCalled = true
			savedPath = path
			if cfg.Metadata.Name != "test-init" {
				t.Errorf("expected config name to be 'test-init', got '%s'", cfg.Metadata.Name)
			}
			return nil
		},
	}

	engine := NewEngine(nil, mockConfigMgr, nil)
	cfg := &types.ClusterConfig{Metadata: types.Metadata{Name: "test-init"}}

	err := engine.Init(context.Background(), cfg)
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	if !saveCalled {
		t.Errorf("expected config manager Save to be called")
	}
	if savedPath != "cluster.yaml" {
		t.Errorf("expected to save to 'cluster.yaml', got '%s'", savedPath)
	}
}

func TestEngineUnimplementedMethods(t *testing.T) {
	engine := NewEngine(nil, nil, nil)
	cfg := &types.ClusterConfig{}
	ctx := context.Background()

	testCases := []struct {
		name string
		err  error
	}{
		{"Failover", engine.Failover(ctx, cfg, "")},
		{"Upgrade", engine.Upgrade(ctx, cfg, "")},
		{"ReplaceNode", engine.ReplaceNode(ctx, cfg, "", "")},
		{"Backup", engine.Backup(ctx, cfg, "")},
		{"Restore", engine.Restore(ctx, cfg, "")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				t.Fatalf("expected error, got nil")
			}
			if tc.err.Error() != "not implemented" {
				t.Fatalf("expected 'not implemented' error, got '%v'", tc.err)
			}
		})
	}
}

//Personal.AI order the ending
