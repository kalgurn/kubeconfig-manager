package kubeconfig

import (
	"log"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Load(kubeConfigPath string) *api.Config {
	kubeConfig, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	return kubeConfig
}
