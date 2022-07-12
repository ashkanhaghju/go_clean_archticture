#!/bin/bash
set -e

readonly service="$1"

protoc --go_out=plugins=grpc:./internal/transport/rpc/grpc/genproto/user/ \
  --proto_path=api/protobuf "api/protobuf/$service.proto"
