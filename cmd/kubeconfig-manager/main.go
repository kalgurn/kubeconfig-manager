package main

import (
	"os"

	"github.com/kalgurn/kubeconfig-manager/internal/command"
)

func main() {

	command.Run(os.Args)

}
