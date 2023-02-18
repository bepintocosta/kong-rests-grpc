#!/bin/bash




kubectl apply -f Namespace.yaml
kubectl apply -f IngressClass.yaml
kubectl apply -f Secret.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -f ClusterRoleBinding.yaml
kubectl apply -f Role.yaml
kubectl apply -f RoleBinding.yaml


#-----------Configmaps should go here
kubectl -n kong-public create configmap kong-plugins-jwt-keycloak --from-file=../plugins/jwt-keycloak --dry-run=client -o yaml | kubectl apply -f -
kubectl -n kong-public create configmap kong-plugins-jwt-keycloak-validators --from-file=../plugins/jwt-keycloak/validators --dry-run=client -o yaml | kubectl apply -f -



kubectl apply -f Deployment.yaml
kubectl apply -f Services.yaml
kubectl apply -f Job.yaml
sed -i "s^{{KEYCLOAK_ISS_URL}}^https://bepintocosta.eu.auth0.com/^g" KongClusterplugin-jwt-keycloak.yaml
kubectl apply -f KongClusterplugin-jwt-keycloak.yaml


