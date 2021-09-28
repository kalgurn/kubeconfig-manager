package cmd_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/cmd"
	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd/api"
)

var ctxCmd = &cobra.Command{
	Use:   "ctx",
	Short: "Desc",
	RunE: func(command *cobra.Command, args []string) error {
		return cmd.CtxComposer(command, args)
	},
}

func TestCtx(t *testing.T) {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	cfg.Clusters["ctx1-test"] = api.NewCluster()
	cfg.AuthInfos["ctx1-test"] = api.NewAuthInfo()
	cfg.CurrentContext = "ctx2-test"
	kubeconfig.Export("ctx1-test", cfg)
	os.Setenv("KUBECONFIG", "ctx1-test.yaml")
	output, err := executeCommand(ctxCmd, "ctx1-test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
