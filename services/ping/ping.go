package ping

import (
	"fmt"
	"strings"

	"github.com/emadghaffari/grpc_oauth_service/databases/protos/pingpb"
	"github.com/emadghaffari/res_errors/errors"
)

// Ping func services
func Ping(req *pingpb.PingRequest) (*pingpb.PingResponse, errors.ResError) {
	if strings.TrimSpace(req.GetValue()) == "" {
		return nil, errors.HandlerBadRequest("invalid ping request")
	}
	return &pingpb.PingResponse{
		Result: fmt.Sprintf("Hi %s", strings.TrimSpace(req.GetValue())),
	}, nil
}
