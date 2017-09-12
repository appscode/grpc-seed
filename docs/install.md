> New to Plow? Please start [here](/docs/tutorial.md).

# Installation Guide

## Using YAML
Plow can be installed using YAML files includes in the [/hack/deploy](/hack/deploy) folder.

```console
# Install without RBAC roles
$ curl https://raw.githubusercontent.com/appscode/plow/master/hack/deploy/without-rbac.yaml \
  | kubectl apply -f -


# Install with RBAC roles
$ curl https://raw.githubusercontent.com/appscode/plow/master/hack/deploy/with-rbac.yaml \
  | kubectl apply -f -
```

## Verify installation
To check if Plow operator pods have started, run the following command:
```console
$ kubectl get pods --all-namespaces -l app=plow --watch
```

Once the operator pods are running, you can cancel the above command by typing `Ctrl+C`.
