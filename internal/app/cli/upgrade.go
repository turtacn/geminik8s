package cli

import (
	"github.com/spf13/cobra"
)

// NewUpgradeCmd creates the 'upgrade' command.
func NewUpgradeCmd(appCtx *AppContext) *cobra.Command {
	var version string
	var strategy string

	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade the cluster to a new version",
		Long:  `Performs a controlled upgrade of the geminik8s cluster components to the specified version.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCtx.Logger.Infof("Loading configuration from '%s'", cfgFile)
			cfg, err := appCtx.ConfigManager.Load(cfgFile)
			if err != nil {
				appCtx.Logger.Errorf("Failed to load configuration: %v", err)
				return err
			}

			appCtx.Logger.Infof("Starting upgrade of cluster '%s' to version '%s' using strategy '%s'", cfg.Metadata.Name, version, strategy)
			if err := appCtx.Orchestrator.Upgrade(cmd.Context(), cfg, version); err != nil {
				appCtx.Logger.Errorf("Upgrade failed: %v", err)
				return err
			}

			appCtx.Logger.Infof("Upgrade to version '%s' completed successfully.", version)
			return nil
		},
	}

	cmd.Flags().StringVar(&version, "version", "", "The target version to upgrade to (e.g., 'v1.1.0') (required)")
	cmd.Flags().StringVar(&strategy, "strategy", "rolling", "The upgrade strategy to use (e.g., 'rolling', 'parallel')")
	cmd.MarkFlagRequired("version")

	return cmd
}

//Personal.AI order the ending
