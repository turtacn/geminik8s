package node

import (
	"context"
	"time"

	custom_errors "github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/types"
)

// Node represents the node domain entity.
// It encapsulates the state and business logic of a single node.
type Node struct {
	ID        string
	Config    *types.NodeConfig
	Status    *types.NodeStatus
	HostMeta  *types.HostMeta // Manages its view of the cluster
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Repository defines the interface for node persistence.
// In this project, node state might be read directly from the node's filesystem (HostMeta)
// or from a central DB. This interface abstracts that.
type Repository interface {
	Save(ctx context.Context, node *Node) error
	FindByID(ctx context.Context, id string) (*Node, error)
	FindByIP(ctx context.Context, ip string) (*Node, error)
}

// NewNode creates a new Node entity.
func NewNode(config *types.NodeConfig, hostMeta *types.HostMeta) (*Node, error) {
	if config.IP == "" {
		return nil, custom_errors.New(custom_errors.ValidationError, "node IP cannot be empty")
	}
	return &Node{
		ID:       config.IP, // Using IP as ID for simplicity
		Config:   config,
		HostMeta: hostMeta,
		Status: &types.NodeStatus{
			Status: types.NodeStatusUnknown,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Promote sets the node's role to Leader.
func (n *Node) Promote() error {
	if n.Config.Role == types.RoleLeader {
		return custom_errors.Newf(custom_errors.ValidationError, "node %s is already a leader", n.ID)
	}
	n.Config.Role = types.RoleLeader
	n.HostMeta.MyID.Role = types.RoleLeader
	n.UpdatedAt = time.Now()
	return nil
}

// Demote sets the node's role to Follower.
func (n *Node) Demote() error {
	if n.Config.Role == types.RoleFollower {
		return custom_errors.Newf(custom_errors.ValidationError, "node %s is already a follower", n.ID)
	}
	n.Config.Role = types.RoleFollower
	n.HostMeta.MyID.Role = types.RoleFollower
	n.UpdatedAt = time.Now()
	return nil
}

// UpdateHealth updates the node's health status.
func (n *Node) UpdateHealth(status types.NodeStatusType, message string) {
	n.Status.Status = status
	n.Status.Message = message
	n.Status.LastHeartbeatTime = time.Now()
	n.UpdatedAt = time.Now()
}

// IsHealthy checks if the node is considered healthy.
func (n *Node) IsHealthy() bool {
	return n.Status.Status == types.NodeStatusHealthy
}

//Personal.AI order the ending
