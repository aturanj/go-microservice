package account

import (
	"context"

	log "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
	uuid "github.com/satori/go.uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService function
func NewService(rep Repository, logger log.Logger) Service {

	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {

	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()

	id := uuid.String()

	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	err := s.repository.CreateUser(ctx, user)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "success", nil
}

func (s service) GetUser(ctx context.Context, ID string) (string, error) {

	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(ctx, ID)

	if err != nil {
		level.Error(logger).Log("err", err)

		return "", err
	}

	logger.Log("Get User", ID)

	return email, nil
}
