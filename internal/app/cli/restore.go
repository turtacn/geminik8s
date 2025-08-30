package cli

import (
	"github.com/spf13/cobra"
)

// NewRestoreCmd creates the 'restore' command.
func NewRestoreCmd(appCtx *AppContext) *cobra.Command {
	var source string

	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore the cluster's data from a backup",
		Long:  `Restores the Kubernetes state from a backup file. This is a destructive operation.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCtx.Logger.Infof("Loading configuration from '%s'", cfgFile)
			cfg, err := appCtx.ConfigManager.Load(cfgFile)
			if err != nil {
				appCtx.Logger.Errorf("Failed to load configuration: %v", err)
				return err
			}

			appCtx.Logger.Warnf("This is a destructive operation and will overwrite the current cluster state.")
			appCtx.Logger.Infof("Starting restore of cluster '%s' from '%s'", cfg.Metadata.Name, source)
			if err := appCtx.Orchestrator.Restore(cmd.Context(), cfg, source); err != nil {
				appCtx.Logger.Errorf("Restore failed: %v", err)
				return err
			}

			appCtx.Logger.Infof("Restore completed successfully.")
			return nil
		},
	}

	cmd.Flags().StringVar(&source, "source", "", "The path to the backup file to restore from (required)")
	cmd.MarkFlagRequired("source")

	return cmd
}

//Personal.AI order the ending
