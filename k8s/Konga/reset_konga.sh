#!/bin/bash

kubectl delete -f konga-namespace.yaml
kubectl delete -f konga-postgres-statefulset.yaml
kubectl delete -f konga-postgresql-service.yaml
kubectl delete -f konga-deployment.yaml
kubectl delete -f konga-service.yaml