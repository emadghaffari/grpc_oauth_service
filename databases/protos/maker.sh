#!/bin/sh

# create access token
# protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=import --go_opt=paths=import databases/protos/accesstokenpb/accesstokenpb.proto 
protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative databases/protos/accesstokenpb/accesstokenpb.proto 
protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative databases/protos/pingpb/ping.proto 

# protoc --go_out=plugins=grpc:. databases/protos/accesstokenpb/accesstokenpb.proto 