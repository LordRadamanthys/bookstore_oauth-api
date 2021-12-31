package services

import (
	"github.com/LordRadamanthys/bookstore_oauth-api/domain/users"
	"github.com/LordRadamanthys/bookstore_oauth-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
	// return nil, errors.BadRequestError("error to create")
}
