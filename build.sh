#!/bin/bash
set -e

environment=$1

usage() {
    echo "usage: ./build.sh release | ./build.sh debug"
}

if [ -z $environment ]; then
    echo "Need environment argument"
    usage
    exit 1
fi

if [ $environment = "release" ]; then
    docker build -t sysopsdev/portfolio:latest .
    docker push sysopsdev/portfolio:latest
    source /home/chris/dls/kubeconf.sh
    PORTFOLIO=$(kubectl get pods --selector=app=portfolio | tail -n 1 | awk '{ print $1 }')
    kubectl delete pod $PORTFOLIO
elif [ $environment = "debug" ]; then
    docker build -t sysopsdev/portfolio:test .
    docker run --name portfolio -p 5000:5000 -d sysopsdev/portfolio:test
    read -p "Run validation tests/Smoke tests. Press Enter when complete..." enter
    docker stop portfolio
    docker rm portfolio
    docker image rm sysopsdev/portfolio:test
else 
    usage
    exit 1
fi