package cmd

import (
	"fmt"
	"os"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func DelComposer(cmd *cobra.Command, args []string) error {
	context := args[0]
	err := Del(context)

	return err
}
func Del(ctx string) error {
	kubeconfigPath, err := kubeconfig.GetConfigPath()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defaultCfg := kubeconfig.Load(kubeconfigPath)
	logger = log.NewLogger(Verbose)
	err = kubeconfig.ContextExists(ctx, defaultCfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger.Debug("deleting username %s", defaultCfg.Contexts[ctx].AuthInfo)
	delete(defaultCfg.AuthInfos, defaultCfg.Contexts[ctx].AuthInfo)
	logger.Debug("deleting cluster %s", defaultCfg.Contexts[ctx].Cluster)
	delete(defaultCfg.Clusters, defaultCfg.Contexts[ctx].Cluster)
	logger.Debug("deleting context %s", ctx)
	delete(defaultCfg.Contexts, ctx)
	logger.Debug("saving config to %s", kubeconfigPath)
	kubeconfig.Save(kubeconfigPath, defaultCfg)
	logger.Info("config saved")
	return nil
}

var deleteCmd = &cobra.Command{
	Use:     "delete [context to delete]",
	Aliases: []string{"del"},
	Short:   "delete contexts defined in the config",
	Long:    "delete contexts defined in the config",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return DelComposer(cmd, args)
	},
}
