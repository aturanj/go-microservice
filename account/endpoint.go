package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints is an endpoint struct :)
type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

//MakeEndpoints makes endpoints
func MakeEndpoints(s Service) Endpoints {

	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)

		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.ID)

		return GetUserResponse{Email: email}, err
	}

}
