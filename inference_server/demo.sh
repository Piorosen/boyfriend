#!/bin/bash

mkdir -P ./model_repository/densenet

wget -O model_repository/densenet/model.onnx https://contentmamluswest001.blob.core.windows.net/content/14b2744cf8d6418c87ffddc3f3127242/9502630827244d60a1214f250e3bbca7/08aed7327d694b8dbaee2c97b8d0fcba/densenet121-1.2.onnx
wget -O model_repository/densenet/config.pbtxt https://raw.githubusercontent.com/triton-inference-server/server/main/docs/examples/model_repository/densenet_onnx/config.pbtxt
wget -O model_repository/densenet/densenet_labels.txt httphttps://raw.githubusercontent.com/triton-inference-server/server/main/docs/examples/model_repository/densenet_onnx/densenet_labels.txt

docker compose up -d
# perf_analyzer -m densenet_onnx --concurrency-range 1:4
# https://docs.nvidia.com/deeplearning/triton-inference-server/user-guide/docs/user_guide/performance_tuning.html#overview