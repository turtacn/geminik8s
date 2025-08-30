package cli

import (
	"github.com/spf13/cobra"
)

// NewBackupCmd creates the 'backup' command.
func NewBackupCmd(appCtx *AppContext) *cobra.Command {
	var destination string

	cmd := &cobra.Command{
		Use:   "backup",
		Short: "Backup the cluster's data",
		Long:  `Performs a backup of the PostgreSQL database, which contains all Kubernetes state.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCtx.Logger.Infof("Loading configuration from '%s'", cfgFile)
			cfg, err := appCtx.ConfigManager.Load(cfgFile)
			if err != nil {
				appCtx.Logger.Errorf("Failed to load configuration: %v", err)
				return err
			}

			appCtx.Logger.Infof("Starting backup of cluster '%s' to '%s'", cfg.Metadata.Name, destination)
			if err := appCtx.Orchestrator.Backup(cmd.Context(), cfg, destination); err != nil {
				appCtx.Logger.Errorf("Backup failed: %v", err)
				return err
			}

			appCtx.Logger.Infof("Backup completed successfully.")
			return nil
		},
	}

	cmd.Flags().StringVar(&destination, "destination", "./backup.sql.gz", "The path to save the backup file")

	return cmd
}

//Personal.AI order the ending
