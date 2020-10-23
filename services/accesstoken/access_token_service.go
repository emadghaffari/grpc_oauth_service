package accesstoken

import (
	"strings"

	"github.com/emadghaffari/grpc_oauth_service/databases/protos/accesstokenpb"
	"github.com/emadghaffari/grpc_oauth_service/model/accesstoken"
	"github.com/emadghaffari/res_errors/errors"
)

// Get func services
func Get(req *accesstokenpb.GetAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError) {
	if strings.TrimSpace(req.GetAccessToken()) == ""{
		return nil, errors.HandlerBadRequest("invalid access_token for get from oauth service")
	}
	return accesstoken.AccessToken.Get(req)
}

// Store func services
func Store(req *accesstokenpb.StoreAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError) {
	if req.GetClientId() == 0  || req.GetUserId() == 0 {
		return nil, errors.HandlerBadRequest("invalid client or user for get from oauth service")
	}
	return accesstoken.AccessToken.Store(req)
}

// Delete func
func Delete(req *accesstokenpb.DeleteAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError) {
	if strings.TrimSpace(req.GetAccessToken()) == ""{
		return nil, errors.HandlerBadRequest("invalid access_token for get from oauth service")
	}
	return accesstoken.AccessToken.Delete(req)
}

// Update func
func Update(req *accesstokenpb.UpdateAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError) {
	if strings.TrimSpace(req.GetAccessToken()) == ""{
		return nil, errors.HandlerBadRequest("invalid access_token for get from oauth service")
	}
	return accesstoken.AccessToken.Update(req)
}