package libs

import (
	"github.com/backent/go-article/controllers/auth"
	"github.com/backent/go-article/controllers/user"
	"github.com/backent/go-article/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	userController user.ControllerUserInterface,
	authController auth.ControllerAuthInterface,
) *httprouter.Router {
	router := httprouter.New()

	router.POST("/login", authController.Login)

	router.POST("/users", userController.Create)
	router.PUT("/users/:id", userController.Update)
	router.DELETE("/users/:id", userController.Delete)
	router.GET("/users/:id", userController.FindById)
	router.GET("/users", userController.FindAll)

	router.PanicHandler = exception.PanicHandler
	return router
}
