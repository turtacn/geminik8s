package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

// NewStatusCmd creates the 'status' command.
func NewStatusCmd(appCtx *AppContext) *cobra.Command {
	var outputFormat string

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of the cluster",
		Long:  `Checks the health and status of the geminik8s cluster components.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCtx.Logger.Infof("Loading configuration from '%s'", cfgFile)
			cfg, err := appCtx.ConfigManager.Load(cfgFile)
			if err != nil {
				appCtx.Logger.Errorf("Failed to load configuration: %v", err)
				return err
			}

			appCtx.Logger.Infof("Checking status for cluster '%s'...", cfg.Metadata.Name)
			status, err := appCtx.Orchestrator.GetStatus(cmd.Context(), cfg)
			if err != nil {
				appCtx.Logger.Errorf("Failed to get cluster status: %v", err)
				return err
			}

			// Create a printable status object
			printableStatus := struct {
				ClusterName string `json:"clusterName"`
				Status      string `json:"status"`
			}{
				ClusterName: cfg.Metadata.Name,
				Status:      string(*status),
			}

			switch outputFormat {
			case "json":
				data, _ := json.MarshalIndent(printableStatus, "", "  ")
				fmt.Println(string(data))
			case "yaml":
				data, _ := yaml.Marshal(printableStatus)
				fmt.Println(string(data))
			default: // table
				fmt.Printf("Cluster: %s\n", printableStatus.ClusterName)
				fmt.Printf("Status:  %s\n", printableStatus.Status)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, json, yaml)")

	return cmd
}

//Personal.AI order the ending
