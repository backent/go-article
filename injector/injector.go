//go:build wireinject
// +build wireinject

package injector

import (
	controllersUser "github.com/backent/go-article/controllers/user"
	"github.com/backent/go-article/libs"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	servicesUser "github.com/backent/go-article/services/user"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var userSet = wire.NewSet(
	repositoriesUser.NewRepositoryMysqlImpl,
	servicesUser.NewServiceUser,
	controllersUser.NewControllerUser,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		userSet,
		libs.Initiate,
		libs.NewValidator,
		libs.NewRouter,
	)
	return nil
}
