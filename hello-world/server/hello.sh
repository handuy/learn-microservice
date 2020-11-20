#!/bin/bash
protoc --proto_path=$HOME/go/src/learn-microservice/hello-world/server/proto:. --micro_out=./proto --go_out=./proto hello.proto
