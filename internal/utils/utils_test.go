package utils_test

import (
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/utils"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestConstructConfig(t *testing.T) {
	c := api.NewConfig()
	c.Contexts["ctx1-test"] = &api.Context{
		Cluster:  "ctx1-test",
		AuthInfo: "ctx1-test",
	}
	c.Clusters["ctx1-test"] = &api.Cluster{
		Server: "test",
	}
	c.AuthInfos["ctx1-test"] = &api.AuthInfo{
		Token: "test",
	}
	c.CurrentContext = "ctx1-test"

	utils.ConstructConfig("ctx1-test", c)
	utils.ConstructConfig("ctx2-test", c)
}
