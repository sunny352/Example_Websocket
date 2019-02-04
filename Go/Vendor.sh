#!/usr/bin/env bash

#echo "set proxy"
#export http_proxy=http://127.0.0.1:1081
#export https_proxy=http://127.0.0.1:1081

echo "set GOPATH"
export GOPATH=${PWD}:${GOPATH}
export PATH=${PWD}/bin:$PATH

echo "vendor Server"
cd src/Server
govendor init
govendor fetch github.com/gorilla/websocket@v1.4.0
cd ../..

echo "vendor Client"
cd src/Client
govendor init
govendor fetch github.com/gorilla/websocket@v1.4.0
cd ../..