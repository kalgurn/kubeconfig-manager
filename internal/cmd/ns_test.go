package cmd_test

import (
	"fmt"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/cmd"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd/api"
)

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Desc",
	RunE: func(command *cobra.Command, args []string) error {
		return cmd.NsComposer(command, args)
	},
}

func TestNs(t *testing.T) {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	cfg.Clusters["ctx1-test"] = api.NewCluster()
	cfg.AuthInfos["ctx1-test"] = api.NewAuthInfo()
	cfg.CurrentContext = "ctx1-test"
	cmd.Add(cfg)

	output, err := executeCommand(nsCmd, "namespace")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
