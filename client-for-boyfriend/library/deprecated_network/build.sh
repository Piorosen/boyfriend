#!/bin/bash

export CGO_ENABLED=1 && \
export GOOS=darwin && \
export GOARCH=amd64 && \
    go build -v -o libnetwork.a -buildmode=c-archive && \
    mv libnetwork.a ../ && \
    mv libnetwork.h ../../include/network
# go build
# -a 