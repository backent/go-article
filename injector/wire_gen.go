// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	article3 "github.com/backent/go-article/controllers/article"
	auth3 "github.com/backent/go-article/controllers/auth"
	user3 "github.com/backent/go-article/controllers/user"
	"github.com/backent/go-article/libs"
	"github.com/backent/go-article/middlewares"
	"github.com/backent/go-article/repositories/article"
	"github.com/backent/go-article/repositories/auth"
	"github.com/backent/go-article/repositories/user"
	article2 "github.com/backent/go-article/services/article"
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
	repositoryAuthInterface := auth.NewRepositoryAuthJWTImpl()
	userMiddleware := middlewares.NewUserMiddleware(validate, repositoryUserInterface, repositoryAuthInterface)
	serviceUserInterface := user2.NewServiceUser(db, repositoryUserInterface, userMiddleware)
	controllerUserInterface := user3.NewControllerUser(serviceUserInterface)
	authMiddleware := middlewares.NewAuthMiddleware(validate, repositoryUserInterface, repositoryAuthInterface)
	serviceAuthInterface := auth2.NewServiceImpl(db, repositoryUserInterface, repositoryAuthInterface, validate, authMiddleware)
	controllerAuthInterface := auth3.NewControllerAuthImpl(serviceAuthInterface)
	repositoryArticleInterface := article.NewRepositoryArticleMysqlImpl()
	articleMiddleware := middlewares.NewArticleMiddleware(validate, repositoryAuthInterface, repositoryArticleInterface, repositoryUserInterface)
	servicesArticleInterface := article2.NewServicesArticleImpl(db, repositoryArticleInterface, articleMiddleware)
	controllerArticleInterface := article3.NewControllerArticleImpl(servicesArticleInterface)
	router := libs.NewRouter(controllerUserInterface, controllerAuthInterface, controllerArticleInterface)
	return router
}

// injector.go:

var userSet = wire.NewSet(user.NewRepositoryMysqlImpl, user2.NewServiceUser, user3.NewControllerUser, middlewares.NewUserMiddleware)

var articleSet = wire.NewSet(article.NewRepositoryArticleMysqlImpl, article2.NewServicesArticleImpl, article3.NewControllerArticleImpl, middlewares.NewArticleMiddleware)

var authSet = wire.NewSet(auth3.NewControllerAuthImpl, auth2.NewServiceImpl, auth.NewRepositoryAuthJWTImpl, middlewares.NewAuthMiddleware)
