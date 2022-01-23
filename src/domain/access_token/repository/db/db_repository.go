package db

import (
	"github.com/LordRadamanthys/bookstore_oauth-api/src/clients/cassandra"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/domain/access_token"
	"github.com/LordRadamanthys/bookstore_oauth-api/utils/errors"
)

type DbRepository interface {
	GetById(id string) (*access_token.AccessToken, *errors.RestErr)
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
		panic(err)
	}
	defer session.Close()
	return nil, errors.InternalServerError("Database not found!")
}
