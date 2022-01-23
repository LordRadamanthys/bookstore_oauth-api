package access_token

import (
	"strings"

	"github.com/LordRadamanthys/bookstore_oauth-api/utils/errors"
)

type Service interface {
	GetById(id string) (*AccessToken, *errors.RestErr)
}

type Repository interface {
	GetById(id string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(id string) (*AccessToken, *errors.RestErr) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.BadRequestError("Invalid access token")
	}
	accessToken, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
