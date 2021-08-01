package kubeconfig

import (
	"fmt"
	"log"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Export(ctx string, cfg *api.Config) {
	config := fmt.Sprintf("%s.yaml", ctx)
	err := clientcmd.WriteToFile(*cfg, config)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Exported as %s", config)
}
