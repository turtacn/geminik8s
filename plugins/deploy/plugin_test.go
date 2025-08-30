package deploy

import (
	"context"
	"testing"

	"github.com/turtacn/geminik8s/pkg/api"
	"github.com/turtacn/geminik8s/pkg/types"
)

func TestDeployPlugin_Name(t *testing.T) {
	p := New()
	if p.Name() != "deploy" {
		t.Errorf("expected plugin name to be 'deploy', got '%s'", p.Name())
	}
}

func TestDeployPlugin_Validate(t *testing.T) {
	p := New()

	t.Run("valid params", func(t *testing.T) {
		params := api.PluginParams{"config": &types.ClusterConfig{}}
		err := p.Validate(params)
		if err != nil {
			t.Errorf("validation should have passed, but got error: %v", err)
		}
	})

	t.Run("missing config", func(t *testing.T) {
		params := api.PluginParams{}
		err := p.Validate(params)
		if err == nil {
			t.Errorf("validation should have failed due to missing config, but it passed")
		}
	})
}

func TestDeployPlugin_Execute(t *testing.T) {
	p := New()
	cfg := &types.ClusterConfig{
		Metadata: types.Metadata{Name: "test-deploy"},
	}
	params := api.PluginParams{"config": cfg}

	result, err := p.Execute(context.Background(), params)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if !result.Success {
		t.Errorf("expected result to be successful")
	}
	expectedMsg := "Cluster 'test-deploy' deployed successfully."
	if result.Message != expectedMsg {
		t.Errorf("expected message '%s', got '%s'", expectedMsg, result.Message)
	}
}

//Personal.AI order the ending
