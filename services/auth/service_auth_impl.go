package auth

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/middlewares"
	"github.com/backent/go-article/models"
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
	*middlewares.AuthMiddleware
}

func NewServiceImpl(db *sql.DB, userRepository user.RepositoryUserInterface, repositoriesAuth repositoriesAuth.RepositoryAuthInterface, validate *validator.Validate, authMiddleware *middlewares.AuthMiddleware) ServiceAuthInterface {
	return &ServiceAuthImpl{
		DB:                      db,
		RepositoryUserInterface: userRepository,
		Validate:                validate,
		RepositoryAuthInterface: repositoriesAuth,
		AuthMiddleware:          authMiddleware,
	}
}

func (implementation *ServiceAuthImpl) Login(ctx context.Context, request webAuth.LoginRequest) webAuth.LoginResponse {
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := implementation.RepositoryUserInterface.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	if !helpers.CheckPassword(request.Password, user.Password) {
		panic(exception.NewUnAuthorized("unauthorized"))
	}

	generatedToken, err := implementation.RepositoryAuthInterface.Issue(strconv.Itoa(user.Id))
	helpers.PanicIfError(err)

	return webAuth.LoginResponse{
		Token: generatedToken,
	}
}

func (implementation *ServiceAuthImpl) Register(ctx context.Context, request webAuth.RegisterRequest) webAuth.RegisterResponse {

	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.AuthMiddleware.Register(ctx, tx, &request)

	hashedPassword, err := helpers.HashPassword(request.Password)
	helpers.PanicIfError(err)

	user := models.User{
		Username: request.Username,
		Name:     request.Name,
		Password: hashedPassword,
	}

	_, err = implementation.RepositoryUserInterface.Create(ctx, tx, user)
	helpers.PanicIfError(err)

	generatedToken, err := implementation.RepositoryAuthInterface.Issue(strconv.Itoa(request.UserId))
	helpers.PanicIfError(err)

	return webAuth.RegisterResponse{
		Token: generatedToken,
	}
}
