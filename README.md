# github-actions-pull-app
An app repository for demonstrating CI(Github Actions)/CD(ArgoCD).

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

## Add repository and create new app via ArgoCD WebUI
TODO: Add repository and create new app via manifests

## Add a secret for access private container registry
```console
cr_pat='xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'
repo_owner='xxxxxxxxxx'
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