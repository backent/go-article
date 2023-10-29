package user

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/models"
	"github.com/backent/go-article/repositories/user"
	webUser "github.com/backent/go-article/web/user"
)

type ServiceUserImpl struct {
	DB             *sql.DB
	userRepository user.RepositoryUserInterface
}

func NewServiceUser(db *sql.DB, userRepository user.RepositoryUserInterface) ServiceUserInterface {
	return &ServiceUserImpl{
		DB:             db,
		userRepository: userRepository,
	}
}

func (implementation *ServiceUserImpl) Create(ctx context.Context, request webUser.UserRequestCreate) webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user := models.User{
		Username: request.Name,
		Name:     request.Name,
	}

	user, err = implementation.userRepository.Create(ctx, tx, user)
	helpers.PanicIfError(err)

	return webUser.UserModelToResponse(user)
}
func (implementation *ServiceUserImpl) Update(ctx context.Context, request webUser.UserRequestUpdate) webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	_, err = implementation.userRepository.FindById(ctx, tx, request.Id)
	helpers.PanicIfError(exception.NewErrorNotFound(err.Error()))

	user := models.User{
		Id:       request.Id,
		Username: request.Username,
		Name:     request.Name,
	}

	user, err = implementation.userRepository.Update(ctx, tx, user)
	helpers.PanicIfError(err)

	return webUser.UserModelToResponse(user)
}
func (implementation *ServiceUserImpl) Delete(ctx context.Context, id int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	_, err = implementation.userRepository.FindById(ctx, tx, id)
	helpers.PanicIfError(exception.NewErrorNotFound(err.Error()))

	err = implementation.userRepository.Delete(ctx, tx, id)
	helpers.PanicIfError(err)
}
func (implementation *ServiceUserImpl) FindById(ctx context.Context, id int) webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := implementation.userRepository.FindById(ctx, tx, id)
	helpers.PanicIfError(exception.NewErrorNotFound(err.Error()))

	return webUser.UserModelToResponse(user)
}
func (implementation *ServiceUserImpl) FindAll(ctx context.Context) []webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	users, err := implementation.userRepository.FindAll(ctx, tx)
	helpers.PanicIfError(err)

	return webUser.UsersModelToResponses(users)
}