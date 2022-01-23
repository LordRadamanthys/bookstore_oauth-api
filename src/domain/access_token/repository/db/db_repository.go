package db

import (
	"fmt"

	"github.com/LordRadamanthys/bookstore_oauth-api/src/clients/cassandra"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/domain/access_token"
	"github.com/LordRadamanthys/bookstore_oauth-api/utils/errors"
)

const (
	queryGetAccesToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreate        = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?"
)

type DbRepository interface {
	GetById(id string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct{}

// port 9042
// Run commands in cqlsh like this:
// docker exec -it cassandra cqlsh
// Stop the service like this:
// docker stop cassandra
// Start it again like this:
// docker start cassandra
func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		fmt.Println(err)
		return nil, errors.InternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccesToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err.Error() == "not found" {

			return nil, errors.NotFoundError("no access token found")
		}
		return nil, errors.InternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(accessToken access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		fmt.Println(err)
		return errors.InternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(queryCreate,
		accessToken.AccessToken,
		accessToken.UserId,
		accessToken.ClientId,
		accessToken.Expires).Exec(); err != nil {
		if err.Error() == "not found" {

			return errors.NotFoundError("no access token found")
		}
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(accessToken access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		fmt.Println(err)
		return errors.InternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(queryUpdateExpires,
		accessToken.Expires,
		accessToken.AccessToken).Exec(); err != nil {
		if err.Error() == "not found" {

			return errors.NotFoundError("no access token found")
		}
	}
	return nil
}
