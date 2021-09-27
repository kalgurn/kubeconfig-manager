package cmd

import (
	"fmt"
	"os"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd/api"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

func AddComposer(cmd *cobra.Command, args []string) error {
	path := args[0]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("no kubeconfig at path %s", path)
	}
	importedConfig := kubeconfig.Load(path)
	err := Add(importedConfig)
	return err

}
func Add(cfg *api.Config) error {
	kubeconfigPath, err := kubeconfig.GetConfigPath()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defaultCfg := kubeconfig.Load(kubeconfigPath)
	logger = log.NewLogger(Verbose)
	logger.Debug("merging config from ", cfg, " with ", kubeconfigPath)
	kubeconfig.Merge(cfg, defaultCfg)
	logger.Debug("saving config to ", kubeconfigPath)
	kubeconfig.Save(kubeconfigPath, defaultCfg)
	logger.Info("config saved")
	return nil
}

var addCmd = &cobra.Command{
	Use:   "add [path to kubeconfig]",
	Short: "adding kubeconfig context, cluster and user from external to current",
	Long:  "adding kubeconfig context, cluster and user from external to current",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return AddComposer(cmd, args)
	},
}
