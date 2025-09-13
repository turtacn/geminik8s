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
			appCtx.Logger.Info("🚀 Kicking off geminik8s deployment...")

			appCtx.Logger.Debugf("Attempting to load configuration from '%s'", cfgFile)
			cfg, err := appCtx.ConfigManager.Load(cfgFile)
			if err != nil {
				appCtx.Logger.Errorf("❌ Failed to load cluster configuration: %v", err)
				appCtx.Logger.Info("Please ensure a valid 'cluster.yaml' exists or use the 'init' command to create one.")
				return err
			}
			appCtx.Logger.Infof("✅ Loaded configuration for cluster: %s", cfg.Metadata.Name)

			appCtx.Logger.Info("🔥 Starting cluster deployment... (This may take a few minutes)")
			// Here you could use a spinner library for better UX
			if err := appCtx.Orchestrator.Deploy(cmd.Context(), cfg); err != nil {
				appCtx.Logger.Errorf("❌ Deployment failed: %v", err)
				appCtx.Logger.Info("Check the logs for more details. You may need to run 'geminik8s cleanup' before retrying.")
				return err
			}

			appCtx.Logger.Infof("✅ Cluster '%s' deployed successfully!", cfg.Metadata.Name)
			appCtx.Logger.Info("You can now check the status of your cluster with: gemin_k8s status")
			return nil
		},
	}
	// Add deployment-specific flags here if needed, e.g., --timeout, --force
	return cmd
}

//Personal.AI order the ending
