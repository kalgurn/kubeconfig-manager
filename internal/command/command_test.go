package command

import (
	"fmt"
	"os"
	"testing"
)

var validPath string

func init() {
	usrHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	validPath = usrHome + "/.kube/config"

}
func TestGetConfigUset(t *testing.T) {
	configPath := getConfig()
	if configPath != validPath {
		t.Errorf("getConfig error | uset KUBECONFIG - expected %v, got %v", validPath, configPath)
	}
}

func TestGetConfigSet(t *testing.T) {
	os.Setenv("KUBECONFIG", validPath)
	configPath := getConfig()
	if configPath != validPath {
		t.Errorf("getConfig error | set KUBECONFIG - expected %v, got %v", validPath, configPath)
	}

}
