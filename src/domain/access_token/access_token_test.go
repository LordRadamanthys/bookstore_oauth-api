package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstats(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24h")
}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpire(), "brand new access token should not be nil")

	assert.Empty(t, at.AccessToken, "new access token should not have defined access token id")

	assert.True(t, at.UserId == 0, "new access token should not have an associeted user id")

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpire(), "empty access token should be expired by default")

	at.Expires = int(time.Now().UTC().Add(3 * time.Hour).Unix())
	assert.False(t, at.IsExpire(), "access token expiring three hours from now should NOT be expired")

}
