package service

import "context"

type AuthService interface {
	Signin(ctx context.Context, input SigninInput) (string, error)
}

type SigninInput struct {
	username string
	password string
}