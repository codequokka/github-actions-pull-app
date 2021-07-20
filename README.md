# github-actions-pull-app
An app repository for demonstrating CI(Github Actions)/CD(ArgoCD).

[![Test and build](https://github.com/codequokka/github-actions-pull-app/actions/workflows/test-build.yml/badge.svg)](https://github.com/codequokka/github-actions-pull-app/actions/workflows/test-build.yml)

[![Docker build and push](https://github.com/codequokka/github-actions-pull-app/actions/workflows/docker-build-publish.yml/badge.svg)](https://github.com/codequokka/github-actions-pull-app/actions/workflows/docker-build-publish.yml)

## Overview
This repository is demostrating simple CI/CD implementaion.
  - CI is defined in this context as test and build code, build docker image and push it to container registry with Github Actions. 
    - Testing and building code is run with [test-build.yml](https://github.com/codequokka/github-actions-pull-app/blob/main/.github/workflows/test-build.yml) Github Actions workflow.
    - Building docker image and pushing it is also run with [docker-build-publish.yml](https://github.com/codequokka/github-actions-pull-app/blob/main/.github/workflows/docker-build-publish.yml) Github Actionos workflow.
  - CD is defined in this context as deploy new version app to K8s cluster when it's docker image is pushed with ArgoCD.
    - When docker image is pushed to container registry, Create pull request to deploy new versoin app to K8s cluster automatically.  
      This is run with [pull_request.yml](https://github.com/codequokka/github-actions-pull-manifest/blob/main/.github/workflows/pull-request.yml) Github Actionos workflow.
    - When pull request is merged into main branch, ArgoCD in the K8s cluster deploys new version app to K8s cluster.
## TODO
- Introducing Kubeval.
- Introducing Conftest.
- Introducing Open Policy Agent/Gatekeeper.

## Install ArgoCD
```console
$ kubectl create namespace argocd
$ kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

## Generate ArgoCD WebUI URL and admin password.
```console
# For single node K8s cluster has no loadbalancers.
$ kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'
$ k8s_nodeip='xxx.xxx.xxx.xxx' # Replace your K8s cluster node ip address.
$ argocd_port=$(kubectl get service -n argocd argocd-server -o json | jq '.spec.ports[]|select(.name == "https")|.nodePort')
$ argocd_admin_password=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
$ echo "https://${k8s_nodeip}:${argocd_port}"
$ echo "$argocd_admin_password"
```

## Create Personal Access Token and set them as Secrets in this repo
- CR_PAT: a value of repo PAT
- REPO_ACCERSS_PAT: a value of write:packages PAT

## Add repository and create new app via ArgoCD WebUI
TODO: Add repository and create new app via manifests

## Add a secret for access private container registry
```console
$ cr_pat='xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'
$ repo_owner='xxxxxxxxxx'
$ echo $cr_pat | docker login ghcr.io -u $repo_owner --password-stdin
$ kubectl create namespace hello-go
$ kubectl create secret generic regcred --from-file=.dockerconfigjson=${HOME}/.docker/config.json --type=kubernetes.io/dockerconfigjson -n hello-go
```

## Access a deploymented app
```console
# For single node K8s cluster has no loadbalancers.
$ k8s_nodeip='xxx.xxx.xxx.xxx' # Replace your K8s cluster node ip address.
$ app_port=$(kubectl get service -n hello-go -o json | jq '.items[].spec.ports[].nodePort')
$ curl ${k8s_nodeip}:${app_port}
```