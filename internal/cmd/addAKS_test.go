package cmd_test

import (
	"os"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/cmd"
	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestAddAKS(t *testing.T) {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	cfg.Clusters["ctx1-test"] = api.NewCluster()
	cfg.AuthInfos["ctx1-test"] = api.NewAuthInfo()
	cfg.CurrentContext = "ctx1-test"
	kubeconfig.Export("ctx1-test", cfg)
	os.Setenv("KUBECONFIG", "ctx1-test.yaml")

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Desc",
		Run:   emptyRun,
	}
	var addAKSCmd = &cobra.Command{
		Use:   "aks",
		Short: "Desc",
		RunE: func(command *cobra.Command, args []string) error {
			return cmd.AddAKSComposer(command, args)
		},
	}
	var (
		AzureRG        string
		AzureCluster   string
		AzureAdminCred bool
	)
	addCmd.AddCommand(addAKSCmd)
	addAKSCmd.Flags().StringVarP(&AzureRG, "resource-group", "r", "", "Resource Group in which cluster is located")
	addAKSCmd.Flags().StringVarP(&AzureCluster, "cluster", "c", "", "Name of a cluster")
	addAKSCmd.Flags().BoolVarP(&AzureAdminCred, "admin", "a", false, "Download a user or admin credentials")
}
