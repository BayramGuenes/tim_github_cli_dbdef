#!/bin/bash

rm ./cmd/logger/main
cd ./cmd/logger
CGO_ENABLED=0 GOOS=linux go build -o main
cd ..
cd ..
rm ./cmd/receiver/main
cd ./cmd/receiver
CGO_ENABLED=0 GOOS=linux go build -o main
cd ..
cd ..
rm ./cmd/repo/main
cd ./cmd/repo
CGO_ENABLED=0 GOOS=linux go build -o main
cd ..
cd ..


