//go:build wireinject
// +build wireinject

package injector

import (
	controllersArticle "github.com/backent/go-article/controllers/article"
	controllersAuth "github.com/backent/go-article/controllers/auth"
	controllersUser "github.com/backent/go-article/controllers/user"
	"github.com/backent/go-article/libs"
	middlewaresArticle "github.com/backent/go-article/middlewares"
	middlewaresUser "github.com/backent/go-article/middlewares"
	repositoriesArticle "github.com/backent/go-article/repositories/article"
	repositoriesAuth "github.com/backent/go-article/repositories/auth"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	servicesArticle "github.com/backent/go-article/services/article"
	servicesAuth "github.com/backent/go-article/services/auth"
	servicesUser "github.com/backent/go-article/services/user"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var userSet = wire.NewSet(
	repositoriesUser.NewRepositoryMysqlImpl,
	servicesUser.NewServiceUser,
	controllersUser.NewControllerUser,
	middlewaresUser.NewUserMiddleware,
)

var articleSet = wire.NewSet(
	repositoriesArticle.NewRepositoryArticleMysqlImpl,
	servicesArticle.NewServicesArticleImpl,
	controllersArticle.NewControllerArticleImpl,
	middlewaresArticle.NewArticleMiddleware,
)

var authSet = wire.NewSet(
	controllersAuth.NewControllerAuthImpl,
	servicesAuth.NewServiceImpl,
	repositoriesAuth.NewRepositoryAuthJWTImpl,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		authSet,
		userSet,
		articleSet,
		libs.Initiate,
		libs.NewValidator,
		libs.NewRouter,
	)
	return nil
}
