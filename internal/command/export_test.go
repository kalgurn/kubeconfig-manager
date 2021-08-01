package command

import (
	"testing"

	"k8s.io/client-go/tools/clientcmd/api"
)

func TestExport(t *testing.T) {
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

	Export("ctx1-test", c)
	Export("ctx3", c)
}
