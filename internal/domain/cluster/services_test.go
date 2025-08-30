package cluster

import (
	"context"
	"testing"

	"github.com/turtacn/geminik8s/pkg/types"
)

// --- Mocks ---

type mockClusterRepo struct {
	SaveFunc     func(ctx context.Context, cluster *Cluster) error
	FindByIDFunc func(ctx context.Context, id string) (*Cluster, error)
}

func (m *mockClusterRepo) Save(ctx context.Context, cluster *Cluster) error {
	return m.SaveFunc(ctx, cluster)
}

func (m *mockClusterRepo) FindByID(ctx context.Context, id string) (*Cluster, error) {
	return m.FindByIDFunc(ctx, id)
}

func (m *mockClusterRepo) FindByName(ctx context.Context, name string) (*Cluster, error) {
	return nil, nil // Not used in these tests
}

type mockNodeService struct {
	InitializeNodeFunc      func(ctx context.Context, nodeIP string) error
	CheckNodeHealthFunc     func(ctx context.Context, nodeIP string) (bool, error)
	PromoteNodeToLeaderFunc func(ctx context.Context, nodeIP string) error
}

func (m *mockNodeService) InitializeNode(ctx context.Context, nodeIP string) error {
	return m.InitializeNodeFunc(ctx, nodeIP)
}
func (m *mockNodeService) CheckNodeHealth(ctx context.Context, nodeIP string) (bool, error) {
	return m.CheckNodeHealthFunc(ctx, nodeIP)
}
func (m *mockNodeService) PromoteNodeToLeader(ctx context.Context, nodeIP string) error {
	return m.PromoteNodeToLeaderFunc(ctx, nodeIP)
}

type mockStorageService struct {
	ConfigureReplicationFunc func(ctx context.Context, leaderIP, followerIP string) error
	IsReplicationHealthyFunc func(ctx context.Context) (bool, error)
	BackupFunc               func(ctx context.Context, destination string) error
	RestoreFunc              func(ctx context.Context, source string) error
}

func (m *mockStorageService) ConfigureReplication(ctx context.Context, leaderIP, followerIP string) error {
	return m.ConfigureReplicationFunc(ctx, leaderIP, followerIP)
}
func (m *mockStorageService) IsReplicationHealthy(ctx context.Context) (bool, error) {
	return m.IsReplicationHealthyFunc(ctx)
}
func (m *mockStorageService) Backup(ctx context.Context, destination string) error {
	return m.BackupFunc(ctx, destination)
}
func (m *mockStorageService) Restore(ctx context.Context, source string) error {
	return m.RestoreFunc(ctx, source)
}

// --- Tests ---

func TestCreateCluster(t *testing.T) {
	mockRepo := &mockClusterRepo{
		SaveFunc: func(ctx context.Context, cluster *Cluster) error {
			return nil
		},
	}
	service := NewService(mockRepo, nil, nil)

	cfg := &types.ClusterConfig{
		Metadata: types.Metadata{Name: "test-cluster"},
	}

	cluster, err := service.CreateCluster(context.Background(), cfg)
	if err != nil {
		t.Fatalf("CreateCluster failed: %v", err)
	}

	if cluster.ID != "test-cluster" {
		t.Errorf("expected cluster ID to be 'test-cluster', got '%s'", cluster.ID)
	}
	if cluster.Status != types.StatusCreating {
		t.Errorf("expected cluster status to be 'Creating', got '%s'", cluster.Status)
	}
}

func TestDeployCluster(t *testing.T) {
	cfg := &types.ClusterConfig{
		Metadata: types.Metadata{Name: "deploy-test"},
		Spec: types.ClusterSpec{
			Nodes: []types.NodeInfo{
				{IP: "1.1.1.1", Role: types.RoleLeader},
				{IP: "2.2.2.2", Role: types.RoleFollower},
			},
		},
	}
	cluster, _ := NewCluster(cfg)

	mockRepo := &mockClusterRepo{
		FindByIDFunc: func(ctx context.Context, id string) (*Cluster, error) {
			return cluster, nil
		},
		SaveFunc: func(ctx context.Context, cluster *Cluster) error {
			return nil
		},
	}

	initLeaderCalled := false
	initFollowerCalled := false
	mockNodeSvc := &mockNodeService{
		InitializeNodeFunc: func(ctx context.Context, nodeIP string) error {
			if nodeIP == "1.1.1.1" {
				initLeaderCalled = true
			}
			if nodeIP == "2.2.2.2" {
				initFollowerCalled = true
			}
			return nil
		},
	}

	replicationConfigured := false
	mockStorageSvc := &mockStorageService{
		ConfigureReplicationFunc: func(ctx context.Context, leaderIP, followerIP string) error {
			replicationConfigured = true
			return nil
		},
	}

	service := NewService(mockRepo, mockNodeSvc, mockStorageSvc)
	err := service.DeployCluster(context.Background(), "deploy-test")
	if err != nil {
		t.Fatalf("DeployCluster failed: %v", err)
	}

	if !initLeaderCalled {
		t.Errorf("expected InitializeNode to be called for leader")
	}
	if !initFollowerCalled {
		t.Errorf("expected InitializeNode to be called for follower")
	}
	if !replicationConfigured {
		t.Errorf("expected ConfigureReplication to be called")
	}
	if cluster.Status != types.StatusRunning {
		t.Errorf("expected cluster status to be updated to Running, got %s", cluster.Status)
	}
}

//Personal.AI order the ending
