#!/usr/bin/env bash

set -e

export GOPATH=$HOME/go-workspace
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:~/go/bin

echo "Deleting all previously generated files"

cd ~

cd workspace/personal/signaref/proto

rm -rf *.pb.go
rm -rf *.pb.gw.go

echo "Generating pb.go and pb.gw.go files"

sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --go_out=plugins=grpc:. user.proto
sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --go_out=plugins=grpc:. vendor.proto
sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --go_out=plugins=grpc:. common.proto
sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --go_out=plugins=grpc:. product.proto


sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --grpc-gateway_out=logtostderr=true:. user.proto
sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --grpc-gateway_out=logtostderr=true:. vendor.proto
sudo protoc --proto_path=. -I/Users/apple/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis --grpc-gateway_out=logtostderr=true:. product.proto

echo "Generated files"