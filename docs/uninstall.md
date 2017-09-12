# Uninstall Plow
Please follow the steps below to uninstall Plow:

1. Delete the various objects created for Plow operator.
```console
$ ./hack/deploy/uninstall.sh
+ kubectl delete deployment -l app=plow -n kube-system
deployment "plow" deleted
+ kubectl delete service -l app=plow -n kube-system
service "plow" deleted
+ kubectl delete serviceaccount -l app=plow -n kube-system
No resources found
+ kubectl delete clusterrolebindings -l app=plow -n kube-system
No resources found
+ kubectl delete clusterrole -l app=plow -n kube-system
No resources found
```

2. Now, wait several seconds for Plow to stop running. To confirm that Plow operator pod(s) have stopped running, run:
```console
$ kubectl get pods --all-namespaces -l app=plow
```
