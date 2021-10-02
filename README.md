# kubeconfig-manager

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=kalgurn_kubeconfig-manager&metric=alert_status)](https://sonarcloud.io/dashboard?id=kalgurn_kubeconfig-manager)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=kalgurn_kubeconfig-manager&metric=coverage)](https://sonarcloud.io/dashboard?id=kalgurn_kubeconfig-manager)

## Download

To start your experience with a kubeconfig-manager you need to download desired version from the [releases page](https://github.com/kalgurn/kubeconfig-manager/releases).
After you've downloaded a package, you can rename it to something you will be comfortable with and move to your /bin folder, e.g.

```bash
mv kubeconfig-manager-darwin-amd64 /usr/bin/kcmanager
```

You are now ready to go

## Usage

By default _kubeconig-manager_ will use a config defined in a __KUBECONFIG__ environment variable. If there is none, it will use the one from the user home dir, e.g. _~/.kube/config_.

```bash
kcmanager is a tool for kubeconfig management

Usage:
  kcmanager [flags]
  kcmanager [command]

Available Commands:
  add         adding kubeconfig context, cluster and user from external to current
  completion  generate the autocompletion script for the specified shell
  ctx         switch contexts to the one defined in the config
  delete      delete contexts defined in the config
  export      Exports context to the yaml file $context.yaml
  help        Help about any command
  list        list contexts defined in the config
  ns          switch default namespace for the current context
  version     outputs version

Flags:
  -h, --help      help for kcmanager
  -v, --verbose   verbose output

Use "kcmanager [command] --help" for more information about a command.
```

### __kcmanager add__

By default, if not specified with a flag the bahaviour of kcmanager is next:

```bash
Usage:
  kcmanager add [path to kubeconfig] [flags]
  kcmanager add [command]

Available Commands:
  aks         adds kubeconfig downloaded from Azure
  rancher     adds kubeconfig downloaded from a specific rancher installation

Flags:
  -h, --help   help for add

Global Flags:
  -v, --verbose   verbose output

Use "kcmanager add [command] --help" for more information about a command.
```

Currently implemented subcommands for __add__ 

#### __rancher__

This subcommand requires additional environment variable RANCHER_TOKEN to be set. 
This allows to download a kubeconfig for a specific cluster listed in your rancher installation and then add it to your current configuration. Example usage:

```bash
Usage:
  kcmanager add rancher --url=[rancher url] --token=[rancher token|| or the env variable] [flags]

Flags:
  -c, --cluster string   URL to a Rancher
  -h, --help             help for rancher
  -t, --token string     token to a Rancher
  -u, --url string       URL to a Rancher

Global Flags:
  -v, --verbose   verbose output
```

#### __aks__

This subcommand requires additional Azure environment variables to be set:
* SUBSCRIPTION_ID - subscription ID where your cluster is located
* TENANT_ID - Azure tenant ID
* CLIENT_ID - client ID for service-principal
* CLIENT_SECRET secret for the service principal

```bash
Usage:
  kcmanager add aks --resource-group=[Azure RG name] --cluster=[cluster name] --admin[true if set]

Flags:
  -a, --admin                   Download a user/admin credentials
  -c, --cluster string          Name of a cluster
  -h, --help                    help for azure
  -r, --resource-group string   Resource Group in which cluster is located

Global Flags:
  -v, --verbose   verbose output
```
