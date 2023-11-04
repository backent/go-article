package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/helpers"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	webUser "github.com/backent/go-article/web/user"
	"github.com/go-playground/validator/v10"
)

type UserMiddleware struct {
	*validator.Validate
	repositoriesUser.RepositoryUserInterface
}

func NewUserMiddleware(validate *validator.Validate, repositoriesUser repositoriesUser.RepositoryUserInterface) *UserMiddleware {
	return &UserMiddleware{
		Validate:                validate,
		RepositoryUserInterface: repositoriesUser,
	}
}

func (implementation *UserMiddleware) Create(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestCreate) {
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	_, err = implementation.RepositoryUserInterface.FindByUsername(ctx, tx, request.Username)
	if err == nil {
		panic(exception.NewBadRequest("user exists."))
	}
}

func (implementation *UserMiddleware) Update(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestUpdate) {
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)
}
