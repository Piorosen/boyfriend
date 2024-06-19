#!/bin/bash

export CGO_ENABLED=1 && \
    go build -v -o libnetwork.a -buildmode=c-archive && \
    mv libnetwork.a ../ && \
    mv libnetwork.h ../../include/network
