package access_token

import (
	"strings"
	"time"

	"github.com/LordRadamanthys/bookstore_oauth-api/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int    `json:"user_id"`
	ClientId    int    `json:"client_id"`
	Expires     int    `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.BadRequestError("inavalid access token id")
	}
	if at.UserId <= 0 {
		return errors.BadRequestError("inavalid user id")
	}
	if at.ClientId <= 0 {
		return errors.BadRequestError("inavalid client id")
	}
	if at.Expires <= 0 {
		return errors.BadRequestError("inavalid expirantion time")
	}
	return nil

}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: int(time.Now().UTC().Add(expirationTime * time.Hour).Unix()),
	}
}

func (at AccessToken) IsExpire() bool {
	return time.Unix(int64(at.Expires), 0).Before(time.Now().UTC())
}
