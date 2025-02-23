#!/bin/bash

protoc -I proto proto/messages/messages.proto \
  --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
