package access_token_service

import (
	"strings"

	"github.com/LordRadamanthys/bookstore_oauth-api/src/domain/access_token"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/domain/users"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/repository/db"
	"github.com/LordRadamanthys/bookstore_oauth-api/src/repository/rest"
	"github.com/LordRadamanthys/bookstore_oauth-api/utils/errors"
)

type Service interface {
	GetById(id string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(restUsersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: restUsersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(access_token_id string) (*access_token.AccessToken, *errors.RestErr) {
	access_token_id = strings.TrimSpace(access_token_id)

	if len(access_token_id) == 0 {
		return nil, errors.BadRequestError("invalid access token id")
	}

	access_token, err := s.dbRepo.GetById(access_token_id)
	if err != nil {
		return nil, err
	}

	return access_token, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	var user = &users.User{}
	var err *errors.RestErr

	switch request.GrantType {
	case access_token.GrantTypeClientCredentials:
		user, err = s.restUsersRepo.LoginUser(request.ClientID, request.ClientSecret)
	case access_token.GrantTypePassword:
		user, err = s.restUsersRepo.LoginUser(request.Username, request.Password)
	default:
		return nil, errors.BadRequestError("invalid grant type")
	}

	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}
func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.UpdateExpirationTime(at)
}
