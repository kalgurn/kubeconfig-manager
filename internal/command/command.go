package command

import (
	"fmt"
	"os"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"github.com/kalgurn/kubeconfig-manager/internal/rancherClient"
)

var kubeConfigPath = "/.kube/config"

func getConfig() string {
	kubeConfigPath, present := os.LookupEnv("KUBECONFIG")
	usrHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	if !present {
		fmt.Printf("KUBECONFIG variable is not set. Using default from %s/.kube/config\n", usrHome)
		kubeConfigPath = usrHome + "/.kube/config"
	}
	return kubeConfigPath
}

func Run(args []string) {
	kubeConfigPath = getConfig()

	cfg := kubeconfig.Load(kubeConfigPath)
	var err error

	switch args[1] {
	case "list":
		List(cfg)
	case "ctx":
		ctx := args[2]
		err = ContextSwitch(ctx, cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		kubeconfig.Save(kubeConfigPath, cfg)
	case "export":
		ctx := args[2]
		err = Export(ctx, cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "delete":
		ctx := args[2]
		err = Delete(ctx, cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		kubeconfig.Save(kubeConfigPath, cfg)
	case "add":
		switch args[2] {
		case "--rancher":
			if len(os.Args) < 4 {
				fmt.Println("please, provide a valid Rancher URL")
				os.Exit(1)
			} else if len(os.Args) < 5 {
				fmt.Println("please, provide a cluster name")
				os.Exit(1)
			}
			URL := args[3]
			cluster := args[4]
			rancherCfg, err := rancherClient.GetRancherConfig(URL, cluster)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			Add(rancherCfg, cfg)
			kubeconfig.Save(kubeConfigPath, cfg)

		default:
			newCfgPath := args[2]
			newCfg := kubeconfig.Load(newCfgPath)
			Add(newCfg, cfg)
			kubeconfig.Save(kubeConfigPath, cfg)
		}

	default:
		fmt.Println("usage")
		os.Exit(1)
	}

}
