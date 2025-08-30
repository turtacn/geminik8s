package cli

import (
	"github.com/spf13/cobra"
	"github.com/turtacn/geminik8s/pkg/types"
)

// NewInitCmd creates the 'init' command.
func NewInitCmd(appCtx *AppContext) *cobra.Command {
	var (
		clusterName string
		node1IP     string
		node2IP     string
		vip         string
		outputDir   string
	)

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new cluster configuration file",
		Long:  `Creates a new cluster.yaml configuration file with the specified parameters.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCtx.Logger.Infof("Initializing new cluster configuration for '%s'", clusterName)

			cfg := &types.ClusterConfig{
				APIVersion: "geminik8s.turtacn.com/v1alpha1",
				Kind:       "ClusterConfig",
				Metadata: types.Metadata{
					Name: clusterName,
				},
				Spec: types.ClusterSpec{
					Network: types.NetworkConfig{
						VIP: vip,
					},
					Nodes: []types.NodeInfo{
						{IP: node1IP, Role: types.RoleLeader},
						{IP: node2IP, Role: types.RoleFollower},
					},
				},
			}

			// The orchestrator's Init method is responsible for saving the file.
			// We could make the path a parameter if needed.
			if err := appCtx.Orchestrator.Init(cmd.Context(), cfg); err != nil {
				appCtx.Logger.Errorf("Failed to initialize configuration: %v", err)
				return err
			}

			appCtx.Logger.Infof("Successfully created cluster configuration 'cluster.yaml'")
			return nil
		},
	}

	cmd.Flags().StringVar(&clusterName, "name", "", "Cluster name (required)")
	cmd.Flags().StringVar(&node1IP, "node1-ip", "", "IP address of the first node (leader) (required)")
	cmd.Flags().StringVar(&node2IP, "node2-ip", "", "IP address of the second node (follower) (required)")
	cmd.Flags().StringVar(&vip, "vip", "", "Virtual IP for the cluster (required)")
	cmd.Flags().StringVar(&outputDir, "config-dir", ".", "Directory to save the configuration file")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("node1-ip")
	cmd.MarkFlagRequired("node2-ip")
	cmd.MarkFlagRequired("vip")

	return cmd
}

//Personal.AI order the ending
