package azureAKSClient

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	yaml "github.com/ghodss/yaml"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	// "gopkg.in/yaml.v2"
)

var (
	tenantID       string
	subscriptionID string
	clientID       string
	clientSecret   string
	resource       string
	resourceURL    string
)

const (
	apiVersion = "2021-05-01"
)

func IfAdmin(cond bool) string {
	if cond {
		return fmt.Sprintf("listCluster%sCredential", "Admin")
	}
	return fmt.Sprintf("listCluster%sCredential", "User")
}

func GetConfig(rg string, name string, admin bool) (*api.Config, error) {
	token := GetToken()
	tokenString := fmt.Sprintf("%s %s", token.Type, token.Token)
	URL := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ContainerService/managedClusters/%s/%s?api-version=%s", resource, subscriptionID, rg, name, IfAdmin(admin), apiVersion)
	req, err := http.NewRequest("POST", URL, nil)
	RespError(err)

	// add authorization header to the req
	req.Header.Add("Authorization", tokenString)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	RespError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	RespError(err)
	var kubeconfigs AzureAKSCredentials
	err = json.Unmarshal([]byte(body), &kubeconfigs)
	if err != nil {
		fmt.Println(err)
	}

	configEncoded, _ := base64.StdEncoding.DecodeString(kubeconfigs.Kubeconfigs[0].Value)
	config, err := yaml.YAMLToJSON(configEncoded)
	if err != nil {
		fmt.Println(err)
	}

	cfg, err := clientcmd.Load([]byte(config))
	RespError(err)

	return cfg, nil
}

func GetToken() Token {
	tenantID = GetOSVar("TENANT_ID")
	clientID = GetOSVar("CLIENT_ID")
	clientSecret = GetOSVar("CLIENT_SECRET")
	subscriptionID = GetOSVar("SUBSCRIPTION_ID")
	resourceURL = fmt.Sprintf("%s/oauth2/token", tenantID)
	resource = "https://management.azure.com"
	URL := "https://login.microsoftonline.com/"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("resource", resource)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	u, _ := url.ParseRequestURI(URL)
	u.Path = resourceURL
	u.RawQuery = data.Encode()
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	RespError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	RespError(err)
	var authData AuthData
	err = json.Unmarshal([]byte(body), &authData)
	if err != nil {
		fmt.Println(err)
	}
	authToken := Token{
		Type:  authData.TokenType,
		Token: authData.AccessToken,
	}
	return authToken
}

func GetOSVar(envVar string) string {
	value, present := os.LookupEnv(envVar)
	if !present {
		err := fmt.Sprintf("environment variable %s not set", envVar)
		RespError(errors.New(err))
		return ""
	}
	return value
}

func RespError(err error) {
	if err != nil {
		fmt.Printf("there was an error during the call execution: %s\n", err)
	}
}
