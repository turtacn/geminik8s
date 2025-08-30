package cluster

import (
	"context"
	"time"

	custom_errors "github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/types"
)

// Cluster represents the cluster domain entity.
// It encapsulates the state and business logic of a cluster.
type Cluster struct {
	ID        string
	Config    *types.ClusterConfig
	Status    types.ClusterStatus
	Nodes     []*Node // Reference to node domain objects
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Node represents a node within the cluster domain.
// This is a simplified view for the cluster aggregate.
type Node struct {
	ID   string
	IP   string
	Role types.NodeRole
}

// Repository defines the interface for cluster persistence.
type Repository interface {
	Save(ctx context.Context, cluster *Cluster) error
	FindByID(ctx context.Context, id string) (*Cluster, error)
	FindByName(ctx context.Context, name string) (*Cluster, error)
}

// NewCluster creates a new Cluster entity from a configuration.
func NewCluster(config *types.ClusterConfig) (*Cluster, error) {
	// Basic validation can happen here
	if config.Metadata.Name == "" {
		return nil, custom_errors.New(custom_errors.ValidationError, "cluster name cannot be empty")
	}

	cluster := &Cluster{
		ID:        config.Metadata.Name, // Using name as ID for simplicity
		Config:    config,
		Status:    types.StatusCreating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	for _, nodeInfo := range config.Spec.Nodes {
		cluster.Nodes = append(cluster.Nodes, &Node{
			ID:   nodeInfo.IP, // Using IP as ID for simplicity
			IP:   nodeInfo.IP,
			Role: nodeInfo.Role,
		})
	}

	return cluster, nil
}

// ChangeStatus updates the cluster's status.
func (c *Cluster) ChangeStatus(newStatus types.ClusterStatus) {
	c.Status = newStatus
	c.UpdatedAt = time.Now()
}

// IsHealthy checks if the cluster is in a healthy, running state.
func (c *Cluster) IsHealthy() bool {
	return c.Status == types.StatusRunning
}

//Personal.AI order the ending
