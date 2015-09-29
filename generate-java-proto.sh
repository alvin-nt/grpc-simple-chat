#!/usr/bin/env bash
# assumes that you have protoc installed
protoc --plugin=protoc-gen-grpc-java=./protoc-gen-grpc-java --java_out=client/java --grpc-java_out=client/java ChatService.proto
