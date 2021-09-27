package cmd

import (
	"fmt"
	"os"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ctxCmd)
}

func CtxComposer(cmd *cobra.Command, args []string) error {
	context := args[0]
	err := Ctx(context)
	return err

}

func Ctx(ctx string) error {
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
	logger.Debug("switching context to ", ctx)
	defaultCfg.CurrentContext = ctx
	logger.Debug("saving config to ", kubeconfigPath)
	kubeconfig.Save(kubeconfigPath, defaultCfg)
	logger.Debug("config saved")
	logger.Info("context set to ", ctx)
	return nil
}

var ctxCmd = &cobra.Command{
	Use:   "ctx [context to switch to]",
	Short: "switch contexts to the one defined in the config",
	Long:  "switch contexts to the one defined in the config",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return CtxComposer(cmd, args)
	},
}
