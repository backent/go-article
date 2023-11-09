package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/repositories/auth"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	webAuth "github.com/backent/go-article/web/auth"
	"github.com/go-playground/validator/v10"
)

type AuthMiddleware struct {
	*validator.Validate
	repositoriesUser.RepositoryUserInterface
	auth.RepositoryAuthInterface
}

func NewAuthMiddleware(validate *validator.Validate, repositoriesUser repositoriesUser.RepositoryUserInterface, repositoryAuth auth.RepositoryAuthInterface) *AuthMiddleware {
	return &AuthMiddleware{
		Validate:                validate,
		RepositoryUserInterface: repositoriesUser,
		RepositoryAuthInterface: repositoryAuth,
	}
}

func (implementation *AuthMiddleware) Register(ctx context.Context, tx *sql.Tx, request *webAuth.RegisterRequest) {
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	user, err := implementation.RepositoryUserInterface.FindByUsername(ctx, tx, request.Username)
	if err == nil {
		panic(exception.NewBadRequest("user exists."))
	}

	request.UserId = user.Id
}
