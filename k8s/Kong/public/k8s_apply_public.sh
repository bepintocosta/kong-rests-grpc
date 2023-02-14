#!/bin/bash

kubectl apply -f Namespace.yaml
kubectl apply -f IngressClass.yaml
kubectl apply -f Secret.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -f ClusterRoleBinding.yaml
kubectl apply -f Role.yaml
kubectl apply -f RoleBinding.yaml
kubectl apply -f Deployment.yaml
kubectl apply -f Services.yaml
kubectl apply -f Job.yaml


