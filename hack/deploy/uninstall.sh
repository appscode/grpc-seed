#!/bin/bash
set -x

kubectl delete deployment -l app=plow -n kube-system
kubectl delete service -l app=plow -n kube-system

# Delete RBAC objects, if --rbac flag was used.
kubectl delete serviceaccount -l app=plow -n kube-system
kubectl delete clusterrolebindings -l app=plow -n kube-system
kubectl delete clusterrole -l app=plow -n kube-system
