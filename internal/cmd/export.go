package cmd

import (
	"fmt"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exportCmd)
}

func ExportComposer(cmd *cobra.Command, args []string) error {
	context := args[0]
	err := Export(context)

	return err
}

func Export(ctx string) error {
	kubeconfigPath, err := kubeconfig.GetConfigPath()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defaultCfg := kubeconfig.Load(kubeconfigPath)
	logger = log.NewLogger(Verbose)

	logger.Debug("constructing config from the context %s", ctx)
	if c, err := kubeconfig.Construct(ctx, defaultCfg); err != nil {
		logger.Error("%s", err)
		return fmt.Errorf("error: %s", err)
	} else {
		logger.Debug("saving config")
		kubeconfig.Export(ctx, c)
	}
	return nil
}

var exportCmd = &cobra.Command{
	Use:   "export [context to export from the config]",
	Short: "Exports context to the yaml file $context.yaml",
	Long:  "Exports context to the yaml file $context.yaml",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return ExportComposer(cmd, args)
	},
}
