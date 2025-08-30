package cluster

import (
	"context"

	"github.com/turtacn/geminik8s/internal/domain/node"
	"github.com/turtacn/geminik8s/internal/domain/storage"
	custom_errors "github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/types"
)

// Service provides cluster-related operations.
type Service struct {
	clusterRepo Repository
	nodeService node.ServiceInterface    // Depends on node service for node operations
	storageSvc  storage.ServiceInterface // Depends on storage service for storage operations
}

// NewService creates a new cluster service.
func NewService(clusterRepo Repository, nodeService node.ServiceInterface, storageSvc storage.ServiceInterface) *Service {
	return &Service{
		clusterRepo: clusterRepo,
		nodeService: nodeService,
		storageSvc:  storageSvc,
	}
}

// CreateCluster creates a new cluster entity based on a config.
func (s *Service) CreateCluster(ctx context.Context, config *types.ClusterConfig) (*Cluster, error) {
	cluster, err := NewCluster(config)
	if err != nil {
		return nil, err
	}

	if err := s.clusterRepo.Save(ctx, cluster); err != nil {
		return nil, custom_errors.Wrap(err, custom_errors.DatabaseError, "failed to save new cluster")
	}

	return cluster, nil
}

// DeployCluster orchestrates the deployment of a cluster.
// This is a high-level business workflow.
func (s *Service) DeployCluster(ctx context.Context, clusterID string) error {
	cluster, err := s.clusterRepo.FindByID(ctx, clusterID)
	if err != nil {
		return custom_errors.Wrapf(err, custom_errors.DatabaseError, "could not find cluster with id %s", clusterID)
	}

	// 1. Identify leader and follower nodes from config
	var leaderNode, followerNode *types.NodeInfo
	for _, n := range cluster.Config.Spec.Nodes {
		nodeInfo := n // Make a copy to avoid pointer issues with the loop variable
		if nodeInfo.Role == types.RoleLeader {
			leaderNode = &nodeInfo
		} else {
			followerNode = &nodeInfo
		}
	}
	if leaderNode == nil || followerNode == nil {
		return custom_errors.New(custom_errors.ValidationError, "cluster config must have one leader and one follower")
	}

	// 2. Initialize leader node
	if err := s.nodeService.InitializeNode(ctx, leaderNode.IP); err != nil {
		return custom_errors.Wrapf(err, custom_errors.OrchestratorError, "failed to initialize leader node %s", leaderNode.IP)
	}

	// 3. Initialize follower node
	if err := s.nodeService.InitializeNode(ctx, followerNode.IP); err != nil {
		return custom_errors.Wrapf(err, custom_errors.OrchestratorError, "failed to initialize follower node %s", followerNode.IP)
	}

	// 4. Configure storage and replication
	if err := s.storageSvc.ConfigureReplication(ctx, leaderNode.IP, followerNode.IP); err != nil {
		return custom_errors.Wrap(err, custom_errors.OrchestratorError, "failed to configure storage replication")
	}

	cluster.ChangeStatus(types.StatusRunning)
	return s.clusterRepo.Save(ctx, cluster)
}

// CheckClusterHealth checks the overall health of the cluster.
func (s *Service) CheckClusterHealth(ctx context.Context, clusterID string) (types.ClusterStatus, error) {
	cluster, err := s.clusterRepo.FindByID(ctx, clusterID)
	if err != nil {
		return types.StatusUnknown, custom_errors.Wrapf(err, custom_errors.DatabaseError, "could not find cluster with id %s", clusterID)
	}

	// Check health of all nodes
	allNodesHealthy := true
	for _, n := range cluster.Nodes {
		healthy, err := s.nodeService.CheckNodeHealth(ctx, n.IP)
		if err != nil || !healthy {
			allNodesHealthy = false
			break
		}
	}

	// Check replication health
	replicationHealthy, err := s.storageSvc.IsReplicationHealthy(ctx)
	if err != nil {
		// log error but don't necessarily fail the whole cluster
	}

	if allNodesHealthy && replicationHealthy {
		cluster.ChangeStatus(types.StatusRunning)
	} else {
		cluster.ChangeStatus(types.StatusDegraded)
	}

	if err := s.clusterRepo.Save(ctx, cluster); err != nil {
		return types.StatusUnknown, err
	}

	return cluster.Status, nil
}

//Personal.AI order the ending
