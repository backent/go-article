package user

import (
	"context"

	webUser "github.com/backent/go-article/web/user"
)

type ServiceUserInterface interface {
	Create(ctx context.Context, request webUser.UserRequestCreate) webUser.UserResponse
	Update(ctx context.Context, request webUser.UserRequestUpdate) webUser.UserResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) webUser.UserResponse
	FindAll(ctx context.Context) []webUser.UserResponse
}
