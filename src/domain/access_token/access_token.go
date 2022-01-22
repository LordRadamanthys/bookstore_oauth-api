package access_token

import (
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int    `json:"user_id`
	ClientId    int    `json:"client_id`
	Expires     int    `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: int(time.Now().UTC().Add(expirationTime * time.Hour).Unix()),
	}
}

func (at AccessToken) IsExpire() bool {
	return time.Unix(int64(at.Expires), 0).Before(time.Now().UTC())
}
