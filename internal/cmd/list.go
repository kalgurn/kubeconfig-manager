package cmd

import (
	"fmt"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)

}

func ListComposer(cmd *cobra.Command, args []string) error {
	err := List()

	return err
}
func List() error {
	kubeconfigPath, err := kubeconfig.GetConfigPath()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defaultCfg := kubeconfig.Load(kubeconfigPath)
	contexts := kubeconfig.Contexts(defaultCfg)
	logger = log.NewLogger(Verbose)
	logger.Debug("Listing contexts from ", kubeconfig.KubeconfigPath)
	for _, k := range kubeconfig.SortedContexts(contexts) {
		if defaultCfg.CurrentContext == k {
			fmt.Println(k, "<-")
		} else {
			fmt.Println(k)
		}
	}
	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list contexts defined in the config",
	Long:  "list contexts defined in the config",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ListComposer(cmd, args)
	},
}
