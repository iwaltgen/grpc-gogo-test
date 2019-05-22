package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the client and server version information",
	Long: `Print the client and server version information.

Cobra application test project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("version called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
