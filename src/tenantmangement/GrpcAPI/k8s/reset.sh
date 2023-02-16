#!/bin/bash

kubectl delete -f tenantmangement-deployment.yaml
kubectl delete -f grpc-http-rules.yml
