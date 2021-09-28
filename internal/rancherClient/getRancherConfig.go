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

func GetToken(token string) string {
	// RANCHER_TOKEN=token-xxxxx:xxxxxxxxxxxxxxxxxxxxxxxxx
	if len(token) <= 1 {
		tokenEnv, present := os.LookupEnv("RANCHER_TOKEN")
		if !present {
			RespError(errors.New("variable RANCHER_TOKEN is not set"))
			os.Exit(1)
		}
		return "Bearer " + tokenEnv
	} else {
		return "Bearer " + token
	}

}
func GetClusterID(URL string, cluster string, token string) string {
	url := fmt.Sprintf("%s/v3/clusters/?name=%s", URL, cluster)
	bearer := GetToken(token)
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
	if len(cluster1.Data) >= 1 {
		return cluster1.Data[0].Id
	}
	return "error"

}

func GetRancherConfig(url string, cluster string, token string) (*api.Config, error) {
	bearer := GetToken(token)
	clusterID := GetClusterID(url, cluster, token)
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
