//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/turtacn/geminik8s/internal/app/config"
	"github.com/turtacn/geminik8s/internal/app/orchestrator"
	"github.com/turtacn/geminik8s/internal/domain/cluster"
	"github.com/turtacn/geminik8s/internal/infrastructure/database"
	"github.com/turtacn/geminik8s/pkg/api"
	"github.com/turtacn/geminik8s/pkg/types"
)

// TestFullDeploymentWorkflowWithDB simulates a full deployment workflow,
// testing the integration between the orchestrator, domain services, and a real database.
func TestFullDeploymentWorkflowWithDB(t *testing.T) {
	ctx := context.Background()

	// 1. Setup a PostgreSQL container
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(wait.DefaultStartupTimeout()),
		),
	)
	if err != nil {
		t.Fatalf("failed to start postgres container: %v", err)
	}
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate postgres container: %v", err)
		}
	}()

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	// 2. Setup dependencies
	db, err := database.NewPostgreSQL(connStr)
	if err != nil {
		t.Fatalf("failed to connect to test postgres: %v", err)
	}
	defer db.Close()

	// Run migrations (we'll create the table manually for this test)
	_, err = db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS clusters (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			vip TEXT,
			status TEXT
		);
	`)
	if err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	clusterRepo := database.NewClusterRepository(db)
	clusterSvc := cluster.NewService(clusterRepo)

	mockPluginMgr := &MockPluginManager{
		ExecuteFunc: func(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
			return &api.PluginResult{Success: true}, nil
		},
	}

	engine := orchestrator.NewEngine(mockPluginMgr, config.NewManager(), clusterSvc)

	// 3. Create a test cluster config
	cfg := &types.ClusterConfig{
		Metadata: types.Metadata{Name: "integration-test-cluster"},
		Spec: types.ClusterSpec{
			Network: types.NetworkConfig{VIP: "10.0.0.1"},
		},
	}

	// 4. Execute
	t.Log("Starting deployment workflow...")
	err = engine.Deploy(context.Background(), cfg)
	if err != nil {
		t.Fatalf("Deploy failed: %v", err)
	}

	// 5. Assert
	// Verify that the cluster was saved to the database by the orchestrator
	var count int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM clusters WHERE name = $1", cfg.Metadata.Name).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query database: %v", err)
	}
	if count != 1 {
		t.Errorf("expected 1 cluster in the database with the test name, found %d", count)
	}
	t.Log("Deployment workflow completed successfully.")
}

// MockPluginManager for the orchestrator
type MockPluginManager struct {
	ExecuteFunc func(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error)
}

func (m *MockPluginManager) Register(plugin api.Plugin) error    { return nil }
func (m *MockPluginManager) Get(name string) (api.Plugin, error) { return nil, nil }
func (m *MockPluginManager) Execute(ctx context.Context, name string, params api.PluginParams) (*api.PluginResult, error) {
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(ctx, name, params)
	}
	return nil, fmt.Errorf("ExecuteFunc not implemented")
}
