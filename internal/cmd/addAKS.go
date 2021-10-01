package cmd

import (
	"fmt"
	"os"

	"github.com/kalgurn/kubeconfig-manager/internal/azureAKSClient"
	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	log "github.com/kalgurn/kubeconfig-manager/internal/logger"
	"github.com/spf13/cobra"
)

var (
	AzureRG        string
	AzureCluster   string
	AzureAdminCred bool
)

//CodeSmell variables to make Sonar happy
const (
	AzureRGString        = "resource-group"
	AzureClusterString   = "cluster"
	AzureAdminCredString = "admin"
)

func init() {

	addAKSCmd.Flags().StringVarP(&AzureRG, AzureRGString, "r", "", "Resource Group in which cluster is located")
	addAKSCmd.Flags().StringVarP(&AzureCluster, AzureClusterString, "c", "", "Name of a cluster")
	addAKSCmd.Flags().BoolVarP(&AzureAdminCred, AzureAdminCredString, "a", false, "Download a user or admin credentials")
	addAKSCmd.MarkFlagRequired(AzureRGString)
	addAKSCmd.MarkFlagRequired(AzureClusterString)
	addCmd.AddCommand(addAKSCmd)
}

func AddAKSComposer(cmd *cobra.Command, args []string) error {
	AzureRG, _ = cmd.Flags().GetString(AzureRGString)
	AzureCluster, _ = cmd.Flags().GetString(AzureClusterString)
	AzureAdminCred, _ = cmd.Flags().GetBool(AzureAdminCredString)

	err := AddAKS(AzureRG, AzureCluster, AzureAdminCred)

	return err
}
func AddAKS(rg string, name string, admin bool) error {
	kubeconfigPath, err := kubeconfig.GetConfigPath()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defaultCfg := kubeconfig.Load(kubeconfigPath)
	logger = log.NewLogger(Verbose)
	logger.Debug("downloading kubeconfig for a cluster ", AzureCluster, " from ", AzureRG)
	azureCfg, err := azureAKSClient.GetConfig(rg, name, admin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger.Debug("merging config with ", kubeconfigPath)
	kubeconfig.Merge(azureCfg, defaultCfg)
	logger.Debug("saving config to ", kubeconfigPath)
	kubeconfig.Save(kubeconfigPath, defaultCfg)
	logger.Info("config saved")

	return nil
}

var addAKSCmd = &cobra.Command{
	Use:   "aks --resource-group=[Azure RG name] --cluster=[cluster name] --admin[true if set]",
	Short: "adds kubeconfig downloaded from Azure",
	Long:  "adds kubeconfig downloaded from Azure",
	RunE: func(cmd *cobra.Command, args []string) error {
		return AddAKSComposer(cmd, args)
	},
}
