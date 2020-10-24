package controllers

import (
	"context"

	"github.com/emadghaffari/grpc_oauth_service/databases/protos/accesstokenpb"
	"github.com/emadghaffari/grpc_oauth_service/services/accesstoken"
)



var(
	// ServerAccessToken var
	ServerAccessToken accessTokenInterface = &accessToken{}
)


// accessTokenInterface interface
// the interface for accessTokens in dao
type accessTokenInterface interface{
	Get(context.Context,*accesstokenpb.GetAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error)
	Store(context.Context,*accesstokenpb.StoreAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error)
	Delete(context.Context, *accesstokenpb.DeleteAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error)
	Update(context.Context, *accesstokenpb.UpdateAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error)
}

// accessToken struct implement all methods in interface
type accessToken struct{}

func (ac *accessToken) Get(ctx context.Context,req *accesstokenpb.GetAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error) {
	response , err := accesstoken.Get(req)
	if err != nil {
		return nil,err
	}
	return response,nil
}

func (ac *accessToken) Store(ctx context.Context,req *accesstokenpb.StoreAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, error){
	response , err := accesstoken.Store(req)
	if err != nil {
		return nil,err
	}
	return response,nil
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