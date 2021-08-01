package rancherClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func getToken() string {
	bearer, present := os.LookupEnv("RANCHER_TOKEN")
	// RANCHER_TOKEN=token-xxxxx:xxxxxxxxxxxxxxxxxxxxxxxxx
	if !present {
		fmt.Println("variable RANCHER_TOKEN is not set")
		os.Exit(1)
	}
	return "Bearer " + bearer
}
func getClusterID(URL string, name string) string {
	url := fmt.Sprintf("%s/v3/clusters/?name=%s", URL, name)
	bearer := getToken()
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error on response.\n[ERROR] -", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error while reading the response bytes:", err)
	}
	var cluster1 clusters
	err = json.Unmarshal([]byte(body), &cluster1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cluster1)
	return cluster1.Data[0].Id

}

func GetRancherConfig(url string, cluster string) (*api.Config, error) {
	bearer := getToken()
	clusterID := getClusterID(url, cluster)
	URL := fmt.Sprintf("%s/v3/clusters/%s?action=generateKubeconfig", url, clusterID)
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		fmt.Println("error on response.\n[ERROR] -", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error while reading the response bytes:", err)
	}
	var config1 config
	err = json.Unmarshal([]byte(body), &config1)
	if err != nil {
		fmt.Println(err)
	}
	cfg, err := clientcmd.Load([]byte(config1.Config))
	if err != nil {
		fmt.Println(err)

	}

	return cfg, nil
}
