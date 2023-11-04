package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/repositories/auth"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	webUser "github.com/backent/go-article/web/user"
	"github.com/go-playground/validator/v10"
)

type UserMiddleware struct {
	*validator.Validate
	repositoriesUser.RepositoryUserInterface
	auth.RepositoryAuthInterface
}

func NewUserMiddleware(validate *validator.Validate, repositoriesUser repositoriesUser.RepositoryUserInterface, repositoryAuth auth.RepositoryAuthInterface) *UserMiddleware {
	return &UserMiddleware{
		Validate:                validate,
		RepositoryUserInterface: repositoriesUser,
		RepositoryAuthInterface: repositoryAuth,
	}
}

func (implementation *UserMiddleware) Create(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestCreate) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	_, err = implementation.RepositoryUserInterface.FindByUsername(ctx, tx, request.Username)
	if err == nil {
		panic(exception.NewBadRequest("user exists."))
	}
}

func (implementation *UserMiddleware) Update(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestUpdate) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)
}

func (implementation *UserMiddleware) Delete(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestDelete) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

}

func (implementation *UserMiddleware) FindById(ctx context.Context) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}

func (implementation *UserMiddleware) FindAll(ctx context.Context) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}
