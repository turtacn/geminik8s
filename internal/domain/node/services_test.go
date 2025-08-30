package node

import (
	"context"
	"os"
	"testing"

	"github.com/turtacn/geminik8s/pkg/types"
)

// --- Mocks ---

type mockNodeRepo struct {
	SaveFunc     func(ctx context.Context, node *Node) error
	FindByIPFunc func(ctx context.Context, ip string) (*Node, error)
}

func (m *mockNodeRepo) Save(ctx context.Context, node *Node) error             { return m.SaveFunc(ctx, node) }
func (m *mockNodeRepo) FindByID(ctx context.Context, id string) (*Node, error) { return nil, nil }
func (m *mockNodeRepo) FindByIP(ctx context.Context, ip string) (*Node, error) {
	return m.FindByIPFunc(ctx, ip)
}

type mockSystemOperator struct {
	RunCommandFunc func(command string, args ...string) (string, error)
}

func (m *mockSystemOperator) RunCommand(command string, args ...string) (string, error) {
	return m.RunCommandFunc(command, args...)
}
func (m *mockSystemOperator) WriteFile(path string, content []byte, perm os.FileMode) error {
	return nil
}
func (m *mockSystemOperator) ReadFile(path string) ([]byte, error) { return nil, nil }

// --- Tests ---

func TestPromoteNodeToLeader(t *testing.T) {
	nodeToPromote := &Node{
		ID: "1.2.3.4",
		Config: &types.NodeConfig{
			Role: types.RoleFollower,
		},
		HostMeta: &types.HostMeta{
			MyID: types.NodeIdentity{Role: types.RoleFollower},
		},
	}

	saveCalled := false
	mockRepo := &mockNodeRepo{
		FindByIPFunc: func(ctx context.Context, ip string) (*Node, error) {
			return nodeToPromote, nil
		},
		SaveFunc: func(ctx context.Context, node *Node) error {
			saveCalled = true
			return nil
		},
	}

	service := NewService(mockRepo, nil, nil)
	err := service.PromoteNodeToLeader(context.Background(), "1.2.3.4")
	if err != nil {
		t.Fatalf("PromoteNodeToLeader failed: %v", err)
	}

	if !saveCalled {
		t.Errorf("expected repository Save to be called")
	}
	if nodeToPromote.Config.Role != types.RoleLeader {
		t.Errorf("expected node role to be updated to Leader, got %s", nodeToPromote.Config.Role)
	}
}

func TestInitializeNode(t *testing.T) {
	commandCalled := false
	mockSysOp := &mockSystemOperator{
		RunCommandFunc: func(command string, args ...string) (string, error) {
			commandCalled = true
			return "ok", nil
		},
	}

	service := NewService(nil, mockSysOp, nil)
	err := service.InitializeNode(context.Background(), "1.2.3.4")
	if err != nil {
		t.Fatalf("InitializeNode failed: %v", err)
	}

	if !commandCalled {
		t.Errorf("expected RunCommand to be called on system operator")
	}
}

//Personal.AI order the ending
