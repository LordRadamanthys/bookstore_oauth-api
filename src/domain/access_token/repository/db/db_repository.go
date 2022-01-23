package db

import (
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

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.InternalServerError("Database not found!")
}
