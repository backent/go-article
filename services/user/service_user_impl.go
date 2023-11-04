package user

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/middlewares"
	"github.com/backent/go-article/models"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	webUser "github.com/backent/go-article/web/user"
)

type ServiceUserImpl struct {
	DB              *sql.DB
	userRepository  repositoriesUser.RepositoryUserInterface
	middlewaresUser *middlewares.UserMiddleware
}

func NewServiceUser(
	db *sql.DB,
	userRepository repositoriesUser.RepositoryUserInterface,
	middlewaresUser *middlewares.UserMiddleware,
) ServiceUserInterface {
	return &ServiceUserImpl{
		DB:              db,
		userRepository:  userRepository,
		middlewaresUser: middlewaresUser,
	}
}

func (implementation *ServiceUserImpl) Create(ctx context.Context, request webUser.UserRequestCreate) webUser.UserResponse {

	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.middlewaresUser.Create(ctx, tx, &request)

	hashedPassword, err := helpers.HashPassword(request.Password)
	helpers.PanicIfError(err)

	user := models.User{
		Username: request.Username,
		Name:     request.Name,
		Password: hashedPassword,
	}

	user, err = implementation.userRepository.Create(ctx, tx, user)
	helpers.PanicIfError(err)

	return webUser.UserModelToResponse(user)
}
func (implementation *ServiceUserImpl) Update(ctx context.Context, request webUser.UserRequestUpdate) webUser.UserResponse {

	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.middlewaresUser.Update(ctx, tx, &request)

	hashedUserPassword := request.SavedPassword
	if request.Password != "" {
		hashedUserPassword, err = helpers.HashPassword(request.Password)
		helpers.PanicIfError(err)
	}

	user := models.User{
		Id:       request.Id,
		Username: request.Username,
		Name:     request.Name,
		Password: hashedUserPassword,
	}

	user, err = implementation.userRepository.Update(ctx, tx, user)
	helpers.PanicIfError(err)

	return webUser.UserModelToResponse(user)
}
func (implementation *ServiceUserImpl) Delete(ctx context.Context, request webUser.UserRequestDelete) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.middlewaresUser.Delete(ctx, tx, &request)

	err = implementation.userRepository.Delete(ctx, tx, request.Id)
	helpers.PanicIfError(err)
}
func (implementation *ServiceUserImpl) FindById(ctx context.Context, id int) webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.middlewaresUser.FindById(ctx)
	user, err := implementation.userRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	return webUser.UserModelToResponse(user)
}
func (implementation *ServiceUserImpl) FindAll(ctx context.Context) []webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.middlewaresUser.FindAll(ctx)

	users, err := implementation.userRepository.FindAll(ctx, tx)
	helpers.PanicIfError(err)

	return webUser.UsersModelToResponses(users)
}
