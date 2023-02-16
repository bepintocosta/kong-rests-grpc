#!/bin/bash

kubectl delete -f Namespace.yaml
kubectl delete -f IngressClass.yaml
kubectl delete -f Secret.yaml
kubectl delete -f ServiceAccount.yaml
kubectl delete -f ClusterRoleBinding.yaml
kubectl delete -f Role.yaml
kubectl delete -f RoleBinding.yaml
kubectl delete -f Deployment.yaml
kubectl delete -f Services.yaml
kubectl delete -f Job.yaml


