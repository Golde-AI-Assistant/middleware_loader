package service

import (
	"context"
	"fmt"

	"middleware_loader/core/domain/enums"
)

type AuthService interface {
	Signin(ctx context.Context, input SigninInput) (string, error)
}

type SigninInput struct {
	username string
	password string
}

func (in SigninInput) Validate() error {
	if len(in.password) < 1 {
		return fmt.Errorf("%w: password is required", enums.ErrValidation)
	}

	return nil
}