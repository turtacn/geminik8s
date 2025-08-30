package cli

import (
	"github.com/spf13/cobra"
)

// NewDeployCmd creates the 'deploy' command.
func NewDeployCmd(appCtx *AppContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a new cluster from a configuration file",
		Long:  `Deploys a geminik8s cluster based on the provided cluster.yaml file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCtx.Logger.Infof("Loading configuration from '%s'", cfgFile)
			cfg, err := appCtx.ConfigManager.Load(cfgFile)
			if err != nil {
				appCtx.Logger.Errorf("Failed to load configuration: %v", err)
				return err
			}

			appCtx.Logger.Infof("Starting deployment for cluster '%s'...", cfg.Metadata.Name)
			if err := appCtx.Orchestrator.Deploy(cmd.Context(), cfg); err != nil {
				appCtx.Logger.Errorf("Deployment failed: %v", err)
				return err
			}

			appCtx.Logger.Infof("Cluster '%s' deployed successfully.", cfg.Metadata.Name)
			return nil
		},
	}
	// Add deployment-specific flags here if needed, e.g., --timeout, --force
	return cmd
}

//Personal.AI order the ending
