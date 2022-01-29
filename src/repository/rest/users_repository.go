package rest

import (
	"encoding/json"
	"time"

	"github.com/LordRadamanthys/bookstore_oauth-api/src/domain/users"
	"github.com/LordRadamanthys/bookstore_oauth-api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRestRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, passowrd string) (*users.User, *errors.RestErr) {

	request := users.UserLoginRequest{
		Email:    email,
		Password: passowrd,
	}
	response := usersRestClient.Post("users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("invalid rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.InternalServerError("invalid error interface")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.InternalServerError("error trying to parse rest response")
	}

	return &user, nil
}
