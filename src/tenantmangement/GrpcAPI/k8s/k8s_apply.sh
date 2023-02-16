#!/bin/bash

kubectl apply -f tenantmangement-deployment.yaml
kubectl apply -f grpc-http-rules.yml
