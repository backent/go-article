package auth

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/helpers"
	repositoriesAuth "github.com/backent/go-article/repositories/auth"
	"github.com/backent/go-article/repositories/user"
	webAuth "github.com/backent/go-article/web/auth"
	"github.com/go-playground/validator/v10"
)

type ServiceAuthImpl struct {
	DB *sql.DB
	user.RepositoryUserInterface
	repositoriesAuth.RepositoryAuthInterface
	*validator.Validate
	secretKeys []byte
}

func NewServiceImpl(db *sql.DB, userRepository user.RepositoryUserInterface, repositoriesAuth repositoriesAuth.RepositoryAuthInterface, validate *validator.Validate) ServiceAuthInterface {
	return &ServiceAuthImpl{
		DB:                      db,
		RepositoryUserInterface: userRepository,
		Validate:                validate,
		secretKeys:              []byte("amfae73j90fj#%&2"),
		RepositoryAuthInterface: repositoriesAuth,
	}
}

func (implementation *ServiceAuthImpl) Login(ctx context.Context, request webAuth.LoginRequest) webAuth.LoginResponse {
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)

	user, err := implementation.RepositoryUserInterface.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	if user.Name != request.Password {
		panic(exception.NewUnAuthorized("unauthorized"))
	}

	generatedToken, err := implementation.RepositoryAuthInterface.Issue(strconv.Itoa(user.Id))
	helpers.PanicIfError(err)

	return webAuth.LoginResponse{
		Token: generatedToken,
	}
}
