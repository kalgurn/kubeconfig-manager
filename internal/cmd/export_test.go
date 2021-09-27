package cmd_test

import (
	"fmt"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/cmd"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Desc",
	RunE: func(command *cobra.Command, args []string) error {
		return cmd.ExportComposer(command, args)
	},
}

func TestExport(t *testing.T) {
	output, err := executeCommand(exportCmd, "ctx1-test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
