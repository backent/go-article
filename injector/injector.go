//go:build wireinject
// +build wireinject

package injector

import (
	controllersAuth "github.com/backent/go-article/controllers/auth"
	controllersUser "github.com/backent/go-article/controllers/user"
	"github.com/backent/go-article/libs"
	middlewaresUser "github.com/backent/go-article/middlewares"
	repositoriesAuth "github.com/backent/go-article/repositories/auth"
	repositoriesUser "github.com/backent/go-article/repositories/user"
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

var authSet = wire.NewSet(
	controllersAuth.NewControllerAuthImpl,
	servicesAuth.NewServiceImpl,
	repositoriesAuth.NewRepositoryAuthJWTImpl,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		authSet,
		userSet,
		libs.Initiate,
		libs.NewValidator,
		libs.NewRouter,
	)
	return nil
}
