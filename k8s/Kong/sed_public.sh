#!/bin/bash



sed -e "s^{{KONG_CONTEXT}}^public^g" ClusterRoleBinding.yaml  > public/ClusterRoleBinding.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" Deployment.yaml  > public/Deployment.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" IngressClass.yaml  > public/IngressClass.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" Job.yaml  > public/Job.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" Namespace.yaml  > public/Namespace.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" Role.yaml  > public/Role.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" RoleBinding.yaml  > public/RoleBinding.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" Secret.yaml  > public/Secret.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" ServiceAccount.yaml  > public/ServiceAccount.yaml
sed -e "s^{{KONG_CONTEXT}}^public^g" Services.yaml  > public/Services.yaml


