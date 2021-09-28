package cmd

import (
	"fmt"
	"os"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/kalgurn/kubeconfig-manager/internal/rancherClient"
	"github.com/spf13/cobra"
)

var (
	URL     string
	Token   string
	Cluster string
)

func init() {
	addRancherCmd.Flags().StringVarP(&URL, "url", "u", "", "URL to a Rancher")
	addRancherCmd.Flags().StringVarP(&Cluster, "cluster", "c", "", "URL to a Rancher")
	addRancherCmd.Flags().StringVarP(&Token, "token", "t", "", "token to a Rancher")
	addRancherCmd.MarkFlagRequired("url")
	addRancherCmd.MarkFlagRequired("cluster")
	addCmd.AddCommand(addRancherCmd)
}

func AddRancherComposer(cmd *cobra.Command, args []string) error {
	URL, _ = cmd.Flags().GetString("url")
	Token, _ = cmd.Flags().GetString("token")
	Cluster, _ = cmd.Flags().GetString("cluter")

	err := AddRancher(URL, Cluster, Token)

	return err
}
func AddRancher(url string, cluster string, token string) error {
	kubeconfigPath, err := kubeconfig.GetConfigPath()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defaultCfg := kubeconfig.Load(kubeconfigPath)
	logger = log.NewLogger(Verbose)
	logger.Debug("downloading kubeconfig for a cluster", Cluster, "from", URL)
	rancherCfg, err := rancherClient.GetRancherConfig(url, cluster, token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger.Debug("merging config with ", kubeconfigPath)
	kubeconfig.Merge(rancherCfg, defaultCfg)
	logger.Debug("saving config to ", kubeconfigPath)
	kubeconfig.Save(kubeconfigPath, defaultCfg)
	logger.Info("config saved")
	return nil
}

var addRancherCmd = &cobra.Command{
	Use:   "rancher --url=[rancher url] --token=[rancher token]",
	Short: "adding kubeconfig downloaded from a specific rancher installation",
	Long:  "adding kubeconfig downloaded from a specific rancher installation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return AddRancherComposer(cmd, args)
	},
}
