package rancherClient_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/rancherClient"
	"k8s.io/client-go/tools/clientcmd"
)

var validTokenAnswer = "token:megasecured"
var rancherResponse = `{
	"config": "apiVersion: v1\nkind: Config\nclusters:\n- name: \"cfg-test\"\n  cluster:\n    server: \"https://cfg-test\"\n\nusers:\n- name: \"cfg-test\"\n  user:\n    token: \"kubeconfig-u-test:cfg-test\"\n\n\ncontexts:\n- name: \"cfg-test\"\n  context:\n    user: \"cfg-test\"\n    cluster: \"cfg-test\"\n\ncurrent-context: \"cfg-test\"\n"
	}`
var validID = "test-id"
var getIDResponse = fmt.Sprintf(`{"data": [{"id": "%v"}]}`, validID)

func init() {

	os.Setenv("RANCHER_TOKEN", validTokenAnswer)

}

func TestGetToken(t *testing.T) {
	token := rancherClient.GetToken()
	if token != "Bearer "+validTokenAnswer {
		t.Errorf("getToken failed - expected %v, got %v", validTokenAnswer, token)
	}
}

func TestGetClusterID(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getIDResponse))
	}))
	defer ts.Close()

	clusterID := rancherClient.GetClusterID(ts.URL, "test-cluster")
	if clusterID != validID {
		t.Errorf("getClusterID failed - expected %v, got %v", validID, clusterID)
	}

}

func TestGetRancherConfig(t *testing.T) {
	var config1 rancherClient.Config
	err := json.Unmarshal([]byte(rancherResponse), &config1)
	if err != nil {
		t.Error(err)
	}
	validConfig, err := clientcmd.Load([]byte(config1.Config))
	if err != nil {
		t.Error(err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/v3/clusters/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getIDResponse))
	})
	mux.HandleFunc("/v3/clusters/test-id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(rancherResponse))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	clusterConfig, err := rancherClient.GetRancherConfig(ts.URL, "test-id")
	if err != nil {
		t.Error("error during the configuration read")
	}
	if clusterConfig.CurrentContext != validConfig.CurrentContext {
		t.Errorf("getRancherConfig failed - got %v\n expected %v", clusterConfig, validConfig)
	}
}

func TestRespErr(t *testing.T) {
	err := errors.New("test")
	rancherClient.RespError(err)
}
