package kubernetes

import (
	"context"

	"github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/api"
	"github.com/turtacn/geminik8s/pkg/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/rest"
)

// k8sClient implements the api.K8sClient interface.
type k8sClient struct {
	clientset *kubernetes.Clientset
}

// NewK8sClient creates a new Kubernetes client from a kubeconfig file.
func NewK8sClient(kubeconfigPath string) (api.K8sClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, errors.Wrap(err, errors.KubernetesError, "failed to build config from kubeconfig")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, errors.KubernetesError, "failed to create kubernetes clientset")
	}

	return &k8sClient{clientset: clientset}, nil
}

// GetNodes retrieves a list of nodes from the cluster.
// This is a simplified version; a real one would convert corev1.Node to types.Node.
func (c *k8sClient) GetNodes(ctx context.Context) ([]types.Node, error) {
	corev1Nodes, err := c.clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, errors.KubernetesError, "failed to list nodes")
	}

	// In a real implementation, you would map corev1Nodes.Items to []types.Node
	var nodes []types.Node
	for _, n := range corev1Nodes.Items {
		// Simplified mapping
		nodes = append(nodes, types.Node{
			Config: types.NodeConfig{
				Name: n.Name,
			},
		})
	}

	return nodes, nil
}

// Apply is a placeholder for applying a Kubernetes manifest.
// A real implementation would parse the manifest and use the appropriate client
// (e.g., AppsV1().Deployments().Apply(...))
func (c *k8sClient) Apply(ctx context.Context, manifest []byte) error {
	return errors.New(errors.KubernetesError, "Apply not implemented")
}

// Delete is a placeholder for deleting a resource from a manifest.
func (c *k8sClient) Delete(ctx context.Context, manifest []byte) error {
	return errors.New(errors.KubernetesError, "Delete not implemented")
}

//Personal.AI order the ending
