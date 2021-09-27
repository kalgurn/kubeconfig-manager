package cmd_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/cmd"
	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"github.com/kalgurn/kubeconfig-manager/internal/rancherClient"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestAddRancher(t *testing.T) {
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
	var addRancherCmd = &cobra.Command{
		Use:   "rancher",
		Short: "Desc",
		RunE: func(command *cobra.Command, args []string) error {
			return cmd.AddRancherComposer(command, args)
		},
	}
	var (
		URL     string
		Token   string
		Cluster string
	)
	addCmd.AddCommand(addRancherCmd)
	addRancherCmd.Flags().StringVarP(&URL, "url", "u", "", "URL to a Rancher")
	addRancherCmd.Flags().StringVarP(&Cluster, "cluster", "c", "", "URL to a Rancher")
	addRancherCmd.Flags().StringVarP(&Token, "token", "t", "", "token to a Rancher")

	var (
		validID         = "test-id"
		getIDResponse   = fmt.Sprintf(`{"data": [{"id": "%v"}]}`, validID)
		rancherResponse = `{
		"config": "apiVersion: v1\nkind: Config\nclusters:\n- name: \"cfg-test\"\n  cluster:\n    server: \"https://cfg-test\"\n\nusers:\n- name: \"cfg-test\"\n  user:\n    token: \"kubeconfig-u-test:cfg-test\"\n\n\ncontexts:\n- name: \"cfg-test\"\n  context:\n    user: \"cfg-test\"\n    cluster: \"cfg-test\"\n\ncurrent-context: \"cfg-test\"\n"
		}`
	)

	var config1 rancherClient.Config

	err := json.Unmarshal([]byte(rancherResponse), &config1)
	if err != nil {
		t.Error(err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/v3/clusters/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getIDResponse))
	})
	mux.HandleFunc("/v3/clusters/test-id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(rancherResponse))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()
	URL = fmt.Sprintf("--url=%s", ts.URL)

	output, err := executeCommand(addCmd, "rancher", URL, "--cluster=name", "--token=token")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
