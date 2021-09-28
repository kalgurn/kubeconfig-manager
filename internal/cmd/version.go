package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   = "empty version"
	BuildDate = "empty date"
	CommitSHA = "empty sha"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "outputs version",
	Long:  "outputs version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nVersion: %v\nBuild date: %s\nCommit: %v\n", Version, BuildDate, CommitSHA)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
