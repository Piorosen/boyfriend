#!/bin/bash

cp ./library/audio/cmake-build-debug/libaudio* ./library
cp ./library/audio/include/audio/c_api.h ./library


export CGO_ENABLED=1 && \
export GOOS=darwin && \
export GOARCH=amd64 && \
    go build -v -o main
    