// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	auth3 "github.com/backent/go-article/controllers/auth"
	user3 "github.com/backent/go-article/controllers/user"
	"github.com/backent/go-article/libs"
	"github.com/backent/go-article/repositories/auth"
	"github.com/backent/go-article/repositories/user"
	auth2 "github.com/backent/go-article/services/auth"
	user2 "github.com/backent/go-article/services/user"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

// Injectors from injector.go:

func InitializeRouter() *httprouter.Router {
	db := libs.Initiate()
	repositoryUserInterface := user.NewRepositoryMysqlImpl()
	validate := libs.NewValidator()
	serviceUserInterface := user2.NewServiceUser(db, repositoryUserInterface, validate)
	controllerUserInterface := user3.NewControllerUser(serviceUserInterface)
	repositoryAuthInterface := auth.NewRepositoryAuthJWTImpl()
	serviceAuthInterface := auth2.NewServiceImpl(db, repositoryUserInterface, repositoryAuthInterface, validate)
	controllerAuthInterface := auth3.NewControllerAuthImpl(serviceAuthInterface)
	router := libs.NewRouter(controllerUserInterface, controllerAuthInterface)
	return router
}

// injector.go:

var userSet = wire.NewSet(user.NewRepositoryMysqlImpl, user2.NewServiceUser, user3.NewControllerUser)

var authSet = wire.NewSet(auth3.NewControllerAuthImpl, auth2.NewServiceImpl, auth.NewRepositoryAuthJWTImpl)
