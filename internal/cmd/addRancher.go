package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/kalgurn/kubeconfig-manager/internal/rancherClient"
	"github.com/spf13/cobra"
)

var (
	RancherURL     string
	RancherToken   string
	RancherCluster string
)

func init() {
	addRancherCmd.Flags().StringVarP(&RancherURL, "url", "u", "", "URL to a Rancher")
	addRancherCmd.Flags().StringVarP(&RancherCluster, "cluster", "c", "", "URL to a Rancher")
	addRancherCmd.Flags().StringVarP(&RancherToken, "token", "t", "", "token to a Rancher")
	addRancherCmd.MarkFlagRequired("url")
	addRancherCmd.MarkFlagRequired("cluster")
	addCmd.AddCommand(addRancherCmd)
}

func AddRancherComposer(cmd *cobra.Command, args []string) error {
	RancherURL, _ = cmd.Flags().GetString("url")
	RancherToken, _ = cmd.Flags().GetString("token")
	RancherCluster, _ = cmd.Flags().GetString("cluster")

	err := AddRancher(RancherURL, RancherCluster, RancherToken)

	return err
}
func AddRancher(url string, cluster string, token string) error {
	kubeconfigPath, err := kubeconfig.GetConfigPath()
	url = strings.TrimSuffix(url, "/")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defaultCfg := kubeconfig.Load(kubeconfigPath)
	logger = log.NewLogger(Verbose)
	logger.Debug("downloading kubeconfig for a cluster", RancherCluster, "from", RancherURL)
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
	Use:   "rancher --url=[rancher url] --token=[rancher token] --cluster=[cluster name]",
	Short: "adds kubeconfig downloaded from a specific rancher installation",
	Long:  "adds kubeconfig downloaded from a specific rancher installation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return AddRancherComposer(cmd, args)
	},
}
