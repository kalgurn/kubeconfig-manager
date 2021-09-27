package cmd

import (
	"fmt"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nsCmd)
}
func NsComposer(cmd *cobra.Command, args []string) error {
	namespace := args[0]
	err := Ns(namespace)

	return err
}
func Ns(ns string) error {
	kubeconfigPath, err := kubeconfig.GetConfigPath()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defaultCfg := kubeconfig.Load(kubeconfigPath)
	logger = log.NewLogger(Verbose)
	ctx := defaultCfg.CurrentContext
	logger.Debug("switching namespace to ", ns)
	defaultCfg.Contexts[ctx].Namespace = ns
	logger.Debug("saving config to ", kubeconfigPath)
	kubeconfig.Save(kubeconfigPath, defaultCfg)
	logger.Debug("config saved")
	logger.Info("namespace set to ", ns)
	return nil
}

var nsCmd = &cobra.Command{
	Use:   "ns [context to switch to]",
	Short: "switch default namespace for the current context",
	Long:  "switch default namespace for the current context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return NsComposer(cmd, args)
	},
}
