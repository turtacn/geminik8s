package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// These variables are meant to be set at build time
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
	BuiltBy = "dev"
)

// NewVersionCmd creates the 'version' command.
func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of geminik8s",
		Long:  `All software has versions. This is geminik8s's.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("geminik8s Version: %s\n", Version)
			fmt.Printf("Git Commit: %s\n", Commit)
			fmt.Printf("Build Date: %s\n", Date)
			fmt.Printf("Built By: %s\n", BuiltBy)
		},
	}
	return cmd
}

//Personal.AI order the ending
