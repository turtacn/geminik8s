package node

import (
	"context"

	custom_errors "github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/api"
)

// ServiceInterface defines the public methods of a node service.
type ServiceInterface interface {
	InitializeNode(ctx context.Context, nodeIP string) error
	PromoteNodeToLeader(ctx context.Context, nodeIP string) error
	CheckNodeHealth(ctx context.Context, nodeIP string) (bool, error)
}

// Service provides node-related operations.
// Note the name "Service" to avoid collision with the "Node" model.
type Service struct {
	nodeRepo       Repository
	systemOperator api.SystemOperator // For file operations, etc.
	k8sClient      api.K8sClient      // For interacting with k8s
}

// NewService creates a new node service.
func NewService(repo Repository, systemOp api.SystemOperator, k8sClient api.K8sClient) ServiceInterface {
	return &Service{
		nodeRepo:       repo,
		systemOperator: systemOp,
		k8sClient:      k8sClient,
	}
}

// InitializeNode prepares a node for joining the cluster.
// This involves steps like installing software, configuring services, etc.
func (s *Service) InitializeNode(ctx context.Context, nodeIP string) error {
	// This is a business workflow. The actual implementation of these steps
	// would be in the infrastructure layer, called via interfaces.
	// 1. SSH to the node
	// 2. Install K3s, PostgreSQL, etc.
	// 3. Configure services
	// 4. Start services
	// 5. Create HostMeta file

	_, err := s.systemOperator.RunCommand("echo", "Initializing node "+nodeIP)
	if err != nil {
		return custom_errors.Wrapf(err, custom_errors.OrchestratorError, "failed to run init command on %s", nodeIP)
	}

	// In a real implementation, we would fetch the node, update its status, and save it.
	// node, err := s.nodeRepo.FindByIP(ctx, nodeIP)
	// ...
	// node.UpdateHealth(types.NodeStatusHealthy, "Initialization complete")
	// s.nodeRepo.Save(ctx, node)

	return nil
}

// PromoteNodeToLeader handles the business logic of promoting a follower node.
func (s *Service) PromoteNodeToLeader(ctx context.Context, nodeIP string) error {
	node, err := s.nodeRepo.FindByIP(ctx, nodeIP)
	if err != nil {
		return custom_errors.Wrapf(err, custom_errors.DatabaseError, "could not find node with ip %s", nodeIP)
	}

	if err := node.Promote(); err != nil {
		return err
	}

	// Here you would orchestrate the necessary infrastructure changes:
	// 1. Update HostMeta on both nodes.
	// 2. Instruct PostgreSQL to switch roles.
	// 3. Reconfigure Kine to point to the new primary DB.
	// 4. Manage the VIP failover.

	return s.nodeRepo.Save(ctx, node)
}

// CheckNodeHealth performs a health check on a single node.
func (s *Service) CheckNodeHealth(ctx context.Context, nodeIP string) (bool, error) {
	// A real health check would involve multiple steps:
	// 1. Ping the node.
	// 2. Check if k3s service is running.
	// 3. Check if postgres service is running.
	// 4. Check if k8s API is responsive.

	// For now, we simulate a successful check.
	_, err := s.k8sClient.GetNodes(ctx)
	if err != nil {
		return false, custom_errors.Wrapf(err, custom_errors.KubernetesError, "failed to get nodes from k8s api on %s", nodeIP)
	}

	return true, nil
}

//Personal.AI order the ending
