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
kcmanager list - Lists all of the current contextes
kcmanager ctx context_name - Switch to a context context_name
kcmanager export context_name - Export one context from your configuration as a separate kubeconfig with user and server configuration. It will be named as a context_name.yaml
kcmanager delete context_name - Delete context_name context and related user/server
kcmanager add - will be described below
```

### _kcmanager add_ options

By default, if not specified with a flag the bahaviour of kcmanager is next:

```bash
kcmanager add context_name.yaml - will add all of the contexts and related user/server configurations to your kubeconfig
```

Currently implemented flags

- _--rancher_
This flag requires additional enviroenmtn variable RANCHER_TOKEN to be set. This allows to download a kubeconfig for a specific cluster listed in your rancher installation and then add it to your current configuration. Example usage:

```bash
export RANCHER_TOKEN=token-xxxxx:xxxxxxxxxxxxxxxxxxxxxxxxx
kcmanager add --rancher RANCHER_URL CLUSTER_NAME
```
