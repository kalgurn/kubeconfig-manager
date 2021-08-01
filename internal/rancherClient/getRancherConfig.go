package rancherClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func GetToken() string {
	bearer, present := os.LookupEnv("RANCHER_TOKEN")
	// RANCHER_TOKEN=token-xxxxx:xxxxxxxxxxxxxxxxxxxxxxxxx
	if !present {
		RespError(errors.New("variable RANCHER_TOKEN is not set"))
		os.Exit(1)
	}
	return "Bearer " + bearer
}
func GetClusterID(URL string, name string) string {
	url := fmt.Sprintf("%s/v3/clusters/?name=%s", URL, name)
	bearer := GetToken()
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	RespError(err)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	RespError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	RespError(err)
	var cluster1 Clusters
	err = json.Unmarshal([]byte(body), &cluster1)
	RespError(err)
	return cluster1.Data[0].Id

}

func GetRancherConfig(url string, cluster string) (*api.Config, error) {
	bearer := GetToken()
	clusterID := GetClusterID(url, cluster)
	URL := fmt.Sprintf("%s/v3/clusters/%s?action=generateKubeconfig", url, clusterID)
	req, err := http.NewRequest("POST", URL, nil)
	RespError(err)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	RespError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	RespError(err)
	var config1 Config
	err = json.Unmarshal([]byte(body), &config1)
	if err != nil {
		fmt.Println(err)
	}
	cfg, err := clientcmd.Load([]byte(config1.Config))
	RespError(err)

	return cfg, nil
}

func RespError(err error) {
	if err != nil {
		fmt.Println("error while reading the response bytes:", err)
	}

}
