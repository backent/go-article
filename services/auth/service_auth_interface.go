package auth

import (
	"context"

	"github.com/backent/go-article/web/auth"
)

type ServiceAuthInterface interface {
	Login(ctx context.Context, request auth.LoginRequest) auth.LoginResponse
	Register(ctx context.Context, request auth.RegisterRequest) auth.RegisterResponse
}
