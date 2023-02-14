#!/bin/bash

kubectl apply -f konga-namespace.yaml
kubectl apply -f konga-postgres-statefulset.yaml
kubectl apply -f konga-postgresql-service.yaml
kubectl apply -f konga-deployment.yaml
kubectl apply -f konga-service.yaml