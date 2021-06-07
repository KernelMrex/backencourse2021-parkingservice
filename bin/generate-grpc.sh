#!/usr/bin/env bash

cd api/parkinglistservice || exit 1;

protoc \
  -I. \
  -I/usr/local/include \
  --go_out=plugins=grpc:. \
  parkinglistservice.proto
