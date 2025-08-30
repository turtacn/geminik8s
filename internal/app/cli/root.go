package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/turtacn/geminik8s/internal/app/config"
	"github.com/turtacn/geminik8s/internal/app/orchestrator"
	"github.com/turtacn/geminik8s/internal/pkg/logger"
	"github.com/turtacn/geminik8s/pkg/api"
)

var (
	cfgFile  string
	logLevel string
	logFile  string
)

// AppContext holds the services that are shared across commands.
type AppContext struct {
	Orchestrator  api.Orchestrator
	ConfigManager api.ConfigManager
	Logger        logger.Logger
}

// NewRootCmd creates the root command for gemin_k8s.
func NewRootCmd() *cobra.Command {
	appCtx := &AppContext{}

	cmd := &cobra.Command{
		Use:   "gemin_k8s",
		Short: "geminik8s is a dual-node HA solution for Kubernetes.",
		Long: `A command-line tool to manage geminik8s clusters,
providing cost-effective high availability for Kubernetes.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Initialize logger
			var output = os.Stdout
			if logFile != "" {
				f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
				if err != nil {
					return fmt.Errorf("failed to open log file: %w", err)
				}
				output = f
			}
			appCtx.Logger = logger.NewLogger(logLevel, output, "text")

			// Initialize services
			appCtx.ConfigManager = config.NewManager()
			pluginManager := orchestrator.NewPluginManager()
			// TODO: Register actual plugins here
			appCtx.Orchestrator = orchestrator.NewEngine(pluginManager, appCtx.ConfigManager, nil) // Pass nil for domain services for now

			return nil
		},
	}

	// Global flags
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "cluster.yaml", "config file (default is cluster.yaml)")
	cmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level (debug, info, warn, error)")
	cmd.PersistentFlags().StringVar(&logFile, "log-file", "", "log file path (default is stdout)")

	// Add subcommands
	cmd.AddCommand(NewInitCmd(appCtx))
	cmd.AddCommand(NewDeployCmd(appCtx))
	cmd.AddCommand(NewStatusCmd(appCtx))
	cmd.AddCommand(NewFailoverCmd(appCtx))
	cmd.AddCommand(NewUpgradeCmd(appCtx))
	cmd.AddCommand(NewReplaceNodeCmd(appCtx))
	cmd.AddCommand(NewBackupCmd(appCtx))
	cmd.AddCommand(NewRestoreCmd(appCtx))
	cmd.AddCommand(NewVersionCmd()) // Version doesn't need the context

	return cmd
}

//Personal.AI order the ending
