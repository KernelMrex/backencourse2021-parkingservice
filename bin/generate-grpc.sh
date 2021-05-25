#!/usr/bin/env bash

cd cmd/parkinglistservice/transport/proto || exit 1;

protoc \
  -I. \
  -I/usr/local/include \
  --go_out=plugins=grpc:. \
  parkinglistservice.proto
