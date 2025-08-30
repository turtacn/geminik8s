package cli

import (
	"github.com/spf13/cobra"
)

// NewFailoverCmd creates the 'failover' command.
func NewFailoverCmd(appCtx *AppContext) *cobra.Command {
	var promoteNode string

	cmd := &cobra.Command{
		Use:   "failover",
		Short: "Manually trigger a failover to the follower node",
		Long:  `Initiates a manual failover process, promoting the specified follower node to become the new leader.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCtx.Logger.Infof("Loading configuration from '%s'", cfgFile)
			cfg, err := appCtx.ConfigManager.Load(cfgFile)
			if err != nil {
				appCtx.Logger.Errorf("Failed to load configuration: %v", err)
				return err
			}

			appCtx.Logger.Infof("Initiating failover for cluster '%s', promoting node '%s'", cfg.Metadata.Name, promoteNode)
			if err := appCtx.Orchestrator.Failover(cmd.Context(), cfg, promoteNode); err != nil {
				appCtx.Logger.Errorf("Failover failed: %v", err)
				return err
			}

			appCtx.Logger.Infof("Failover completed successfully. Node '%s' is the new leader.", promoteNode)
			return nil
		},
	}

	cmd.Flags().StringVar(&promoteNode, "promote", "", "The IP or name of the follower node to promote (required)")
	cmd.MarkFlagRequired("promote")

	return cmd
}

//Personal.AI order the ending
