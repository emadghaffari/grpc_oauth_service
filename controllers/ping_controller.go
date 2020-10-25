package controllers

import (
	"context"

	"github.com/emadghaffari/grpc_oauth_service/databases/protos/pingpb"
	"github.com/emadghaffari/grpc_oauth_service/services/ping"
)

var (
	// PingController var
	PingController pingInterface = &pingStruct{}
)

// pingInterface interface
// the interface for pings in dao
type pingInterface interface {
	Ping(context.Context, *pingpb.PingRequest) (*pingpb.PingResponse, error)
}

// ping struct implement all methods in interface
type pingStruct struct{}

func (p *pingStruct) Ping(ctx context.Context, req *pingpb.PingRequest) (*pingpb.PingResponse, error) {
	response , err := ping.Ping(req)
	if err != nil {
		return nil,err
	}
	return response,nil
}
