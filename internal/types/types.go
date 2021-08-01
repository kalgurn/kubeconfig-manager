package types

type Kubeconfig struct {
	ApiVersion     string       `yaml:"apiVersion"`
	CurrentContext string       `yaml:"current-context"`
	Kind           string       `yaml:"kind"`
	Preferences    []preference `yaml:"preferences"`
	Users          []user       `yaml:"users"`
	Clusters       []cluster    `yaml:"clusters"`
}
type cluster struct {
	Cluster string
}
type context struct{}
type user struct{}
type preference struct{}
