#!/bin/bash

DOCKER_USER=wannazjx
CURRENT_DIR=$(cd "$(dirname "$0")";pwd)

set -x

kubectl delete -f $CURRENT_DIR/../configuration/scheduler-config.yaml
# build wanna-scheduler
go build -o wanna-scheduler $CURRENT_DIR/../.
# build docker image
docker build --no-cache -f $CURRENT_DIR/../Dockerfile -t ${DOCKER_USER}/wanna-scheduler:v1 .
rm -rf ./wanna-scheduler

docker push ${DOCKER_USER}/wanna-scheduler:v1

kubectl apply -f $CURRENT_DIR/../configuration/scheduler-config.yaml

# gc
docker rm $(docker ps -a | grep Exit | awk '{print $1}')
docker rmi $(docker images | grep none | awk '{print $3}')