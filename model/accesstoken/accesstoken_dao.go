package accesstoken

import (
	"fmt"

	"github.com/gocql/gocql"

	"github.com/emadghaffari/grpc_oauth_service/databases/cassandra"
	"github.com/emadghaffari/grpc_oauth_service/databases/protos/accesstokenpb"
	"github.com/emadghaffari/grpc_oauth_service/utils/date"
	"github.com/emadghaffari/grpc_oauth_service/utils/rand"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/res_errors/logger"
)

var(
	// AccessToken var
	AccessToken accessTokenInterface = &accessToken{}
	getAccessTokenQuery      = "SELECT access_token, user_id, client_id, expired_at FROM access_tokens WHERE access_token=?;"
	createAccesstokenQuery   = "INSERT INTO access_tokens (access_token, user_id, client_id, expired_at) VALUES (?, ?, ?, ?)"
	updateAccesstokenExpires = "UPDATE access_tokens SET expired_at=? WHERE access_token=?;"
	deleteAccessTokenRequest = "DELETE FROM access_tokens WHERE access_token=?;"
)


// accessTokenInterface interface
// the interface for accessTokens in dao
type accessTokenInterface interface{
	Get(req *accesstokenpb.GetAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError)
	Store(req *accesstokenpb.StoreAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError)
	Delete(req *accesstokenpb.DeleteAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError)
	Update(req *accesstokenpb.UpdateAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError)
}

// accessToken struct implement all methods in interface
type accessToken struct{}

// get method
// for get values from cassandraDB
func (ac *accessToken) Get(req *accesstokenpb.GetAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError) {
	var result accesstokenpb.AccessTokenResponse
	if err := cassandra.GetSesstion().Query(getAccessTokenQuery,req.GetAccessToken()).Scan(&result.AccessToken,&result.UserId,&result.ClientId,&result.ExpiredAt); err != nil{
		if err == gocql.ErrNotFound{
			logger.Error("Error gocql Not Found",err)
			return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in gocql NotFound: %v",err),err)
		}
		logger.Error("Error in Query for get from cassandra",err)
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Query for get from cassandra: %v",err),err)
	}
	return &result,nil
}

// Store method
// store new accesstoken
func (ac *accessToken) Store(req *accesstokenpb.StoreAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError) {
	accessToken := rand.String(50)
	expire := date.GetFutureTime(0,0,30)
	if err := cassandra.GetSesstion().Query(createAccesstokenQuery,accessToken,req.GetUserId(),req.GetClientId(),expire).Exec(); err != nil {
		logger.Error("Error in Store new accessToken",err)
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Store new accessToken: %v",err),err)
	}
	
	return &accesstokenpb.AccessTokenResponse{
		UserId: req.GetUserId(),
		ClientId: req.GetClientId(),
		AccessToken: accessToken,
		ExpiredAt: expire,
		CreatedAt: date.GetNowString(),
	},nil
}

// Update method
// update old accessToken with new expired time
func (ac *accessToken) Update(req *accesstokenpb.UpdateAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError) {
	expire := date.GetFutureTime(0,0,30)
	if err := cassandra.GetSesstion().Query(updateAccesstokenExpires,expire,req.GetAccessToken()).Exec();err !=nil {
		logger.Error("Error in Update accessToken",err)
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Update accessToken: %v",err),err)
	}
	return ac.Get(&accesstokenpb.GetAccessTokenRequest{AccessToken: req.GetAccessToken()})
}

// Delete a access_token
func (ac *accessToken) Delete(req *accesstokenpb.DeleteAccessTokenRequest) (*accesstokenpb.AccessTokenResponse, errors.ResError) {
	if err := cassandra.GetSesstion().Query(deleteAccessTokenRequest,req.GetAccessToken()).Exec();err !=nil {
		logger.Error("Error in Update accessToken",err)
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Update accessToken: %v",err),err)
	}
	return &accesstokenpb.AccessTokenResponse{},nil
}