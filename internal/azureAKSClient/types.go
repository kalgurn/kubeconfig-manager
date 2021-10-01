package azureAKSClient

type AuthData struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	ExtExpiresIn string `json:"ext_expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	AccessToken  string `json:"access_token"`
}

type AzureAKSCredentials struct {
	Kubeconfigs []Kubeconfig `json:"kubeconfigs"`
}

// type AzureAKSConfigYaml struct {
// 	ApiVersion     string                                `yaml:"apiVersion"`
// 	Clusters       map[string]AzureAKSConfigClustersYaml `yaml:"clusters"`
// 	Contexts       map[string]AzureAKSConfigContextsYaml `yaml:"contexts"`
// 	CurrentContext string                                `yaml:"current-context"`
// 	Kind           string                                `yaml:"kind"`
// 	Preferences    map[string]string                     `yaml:"preferences"`
// 	Users          map[string]AzureAKSConfigUsursYaml    `yaml:"users"`
// }

// type AzureAKSConfigClustersYaml struct {
// 	Cluster []AzureAKSConfigClusterYaml `yaml:"cluster"`
// 	Name    string                      `yaml:"name"`
// }
// type AzureAKSConfigClusterYaml struct {
// 	CertificateAuthorityData string `yaml:"certificate-authority-data"`
// 	Server                   string `yaml:"server"`
// }
// type AzureAKSConfigContextsYaml struct {
// 	Context []AzureAKSConfigContextYaml `yaml:"context"`
// 	Name    string                      `yaml:"name"`
// }
// type AzureAKSConfigContextYaml struct {
// 	CLuster string `yaml:"cluster"`
// 	User    string `yaml:"user"`
// }
// type AzureAKSConfigUsursYaml struct {
// 	User []AzureAKSConfigUserYaml `yaml:"user"`
// 	Name string                   `yaml:"name"`
// }
// type AzureAKSConfigUserYaml struct {
// 	ClientCertificateData string `yaml:"client-certificate-data"`
// 	ClientKeyData         string `yaml:"client-key-data"`
// 	Token                 string `yaml:"token"`
// }

type Kubeconfig struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Token struct {
	Token string
	Type  string
}
