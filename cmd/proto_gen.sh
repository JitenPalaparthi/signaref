#!/usr/bin/env bash

echo "Deleting all previously generated files"

cd ~

cd workspace/personal/signaref/proto

rm -rf *.pb.go
rm -rf *.pb.gw.go

echo "Generating pb.go and pb.gw.go files"

sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --go_out=plugins=grpc:. user.proto
sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --go_out=plugins=grpc:. vendor.proto
sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --go_out=plugins=grpc:. common.proto

sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --grpc-gateway_out=logtostderr=true:. user.proto
sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --grpc-gateway_out=logtostderr=true:. vendor.proto

echo "Generated files"