package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
)

var (
	Verbose bool
	logger  *log.StandardLogger
)

func init() {

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

}

var rootCmd = &cobra.Command{
	Use:   "kcmanager",
	Short: "kcmanager is a tool for kubeconfig management",
	Long:  "kcmanager is a tool for kubeconfig management",

	Run: func(cmd *cobra.Command, args []string) {
		logger = log.NewLogger(Verbose)
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
