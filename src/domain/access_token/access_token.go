package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/LordRadamanthys/bookstore_oauth-api/utils/crypto_utils"
	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
)

const (
	expirationTime             = 24
	GrantTypePassword          = "password"
	GrantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int    `json:"user_id"`
	ClientId    int    `json:"client_id"`
	Expires     int    `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() rest_errors.RestErr {
	switch at.GrantType {
	case GrantTypePassword:
		return at.validatePasswordGrantType()
	case GrantTypeClientCredentials:
		return at.validatePasswordGrantType()
	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameter")
	}

}

func (at *AccessToken) Validate() rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}

	return nil

}
func (at *AccessTokenRequest) validatePasswordGrantType() rest_errors.RestErr {
	if strings.TrimSpace(at.Username) == "" {
		return rest_errors.NewBadRequestError("username can not be empty")
	}

	if strings.TrimSpace(at.Password) == "" {
		return rest_errors.NewBadRequestError("password can not be empty")
	}

	return nil
}

func (at *AccessTokenRequest) validateClientGrantType() rest_errors.RestErr {
	if strings.TrimSpace(at.ClientID) == "" {
		return rest_errors.NewBadRequestError("client_id can not be empty")
	}

	if strings.TrimSpace(at.ClientSecret) == "" {
		return rest_errors.NewBadRequestError("client_secret can not be empty")
	}

	return nil
}

func GetNewAccessToken(userId int) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: int(time.Now().UTC().Add(expirationTime * time.Hour).Unix()),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(int64(at.Expires), 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
