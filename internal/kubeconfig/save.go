package kubeconfig

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Save(path string, cfg *api.Config) {
	err := clientcmd.WriteToFile(*cfg, path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Changes saved at", path)
}
