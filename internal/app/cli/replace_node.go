package cli

import (
	"github.com/spf13/cobra"
)

// NewReplaceNodeCmd creates the 'replace-node' command.
func NewReplaceNodeCmd(appCtx *AppContext) *cobra.Command {
	var oldNode string
	var newNodeIP string

	cmd := &cobra.Command{
		Use:   "replace-node",
		Short: "Replace a node in the cluster",
		Long:  `Replaces a specified node with a new one, handling data migration and cluster reconfiguration.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCtx.Logger.Infof("Loading configuration from '%s'", cfgFile)
			cfg, err := appCtx.ConfigManager.Load(cfgFile)
			if err != nil {
				appCtx.Logger.Errorf("Failed to load configuration: %v", err)
				return err
			}

			appCtx.Logger.Infof("Starting replacement of node '%s' with new node at '%s'", oldNode, newNodeIP)
			if err := appCtx.Orchestrator.ReplaceNode(cmd.Context(), cfg, oldNode, newNodeIP); err != nil {
				appCtx.Logger.Errorf("Node replacement failed: %v", err)
				return err
			}

			appCtx.Logger.Infof("Node replacement completed successfully.")
			return nil
		},
	}

	cmd.Flags().StringVar(&oldNode, "old-node", "", "The IP or name of the node to replace (required)")
	cmd.Flags().StringVar(&newNodeIP, "new-node-ip", "", "The IP address of the new node (required)")
	cmd.MarkFlagRequired("old-node")
	cmd.MarkFlagRequired("new-node-ip")

	return cmd
}

//Personal.AI order the ending
