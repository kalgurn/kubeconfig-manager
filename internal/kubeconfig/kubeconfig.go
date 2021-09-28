package kubeconfig

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

var (
	Cfg                 *api.Config
	KubeconfigPath, err = GetConfigPath()
)

func init() {
	if err != nil {
		fmt.Printf("%s ", err)
	}
	Cfg = Load(KubeconfigPath)
}
