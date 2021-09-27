package kubeconfig

import (
	"fmt"
	"os"
)

func GetConfigPath() (string, error) {
	kubeConfigPath, present := os.LookupEnv("KUBECONFIG")
	usrHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	if !present {
		// fmt.Printf("KUBECONFIG variable is not set. Using default from %s/.kube/config\n", usrHome)
		kubeConfigPath = usrHome + "/.kube/config"
		if _, err := os.Stat(kubeConfigPath); os.IsNotExist(err) {
			return "", fmt.Errorf("no kubeconfig at path %s", kubeConfigPath)
		}
	}
	return kubeConfigPath, nil
}
