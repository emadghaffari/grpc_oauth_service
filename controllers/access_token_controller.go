package controllers

import (
	"context"
	"io"

	"github.com/emadghaffari/grpc_oauth_service/databases/protos/accesstokenpb"
	"github.com/emadghaffari/grpc_oauth_service/services/accesstoken"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/res_errors/logger"
)



var(
	// ClientAccessToken var
	ClientAccessToken accessTokenInterface = &accessToken{}
)


// accessTokenInterface interface
// the interface for accessTokens in dao
type accessTokenInterface interface{
	Get(accesstokenpb.AccessToken_GetServer) error
	Store(accesstokenpb.AccessToken_StoreServer) error
	Delete(context.Context, *accesstokenpb.DeleteAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error)
	Update(context.Context, *accesstokenpb.UpdateAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error)
}

// accessToken struct implement all methods in interface
type accessToken struct{}

func (ac *accessToken) Get(stream accesstokenpb.AccessToken_GetServer) error {
	for{
		req,err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			logger.Error("Error in Recv for Get values from AccessToken",err)
			return errors.HandlerInternalServerError("Error in Recv for Get values from AccessToken",err)
		}
		response, err := accesstoken.Get(req)
		if err != nil{
			return err
		}
		stream.Send(response)
	}
}
func (ac *accessToken) Store(stream accesstokenpb.AccessToken_StoreServer) error {
	for{
		req,err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			logger.Error("Error in Recv for Store AccessToken section",err)
			return errors.HandlerInternalServerError("Error in Recv from Store AccessToken section",err)
		}
		response, err := accesstoken.Store(req)
		if err != nil{
			return err
		}
		stream.Send(response)
	}
}
func (ac *accessToken) Delete(ctx context.Context,req *accesstokenpb.DeleteAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error) {
	response , err := accesstoken.Delete(req)
	if err != nil {
		return nil,err
	}
	return response,nil
}
func (ac *accessToken) Update(ctx context.Context,req *accesstokenpb.UpdateAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error) {
	response , err := accesstoken.Update(req)
	if err != nil {
		return nil,err
	}
	return response,nil
}