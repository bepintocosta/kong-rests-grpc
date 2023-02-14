#!/bin/bash



sed -e "s^{{KONG_CONTEXT}}^admin^g" ClusterRoleBinding.yaml  > admin/ClusterRoleBinding.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" Deployment.yaml  > admin/Deployment.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" IngressClass.yaml  > admin/IngressClass.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" Job.yaml  > admin/Job.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" Namespace.yaml  > admin/Namespace.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" Role.yaml  > admin/Role.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" RoleBinding.yaml  > admin/RoleBinding.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" Secret.yaml  > admin/Secret.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" ServiceAccount.yaml  > admin/ServiceAccount.yaml
sed -e "s^{{KONG_CONTEXT}}^admin^g" Services.yaml  > admin/Services.yaml


