package libs

import (
	"github.com/backent/go-article/controllers/article"
	"github.com/backent/go-article/controllers/auth"
	"github.com/backent/go-article/controllers/user"
	"github.com/backent/go-article/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	userController user.ControllerUserInterface,
	authController auth.ControllerAuthInterface,
	articleController article.ControllerArticleInterface,
) *httprouter.Router {
	router := httprouter.New()

	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)

	router.POST("/users", userController.Create)
	router.PUT("/users/:id", userController.Update)
	router.DELETE("/users/:id", userController.Delete)
	router.GET("/users/:id", userController.FindById)
	router.GET("/users", userController.FindAll)

	router.POST("/articles", articleController.Create)
	router.PUT("/articles/:id", articleController.Update)
	router.DELETE("/articles/:id", articleController.Delete)
	router.GET("/articles/:id", articleController.FindById)
	router.GET("/articles", articleController.FindAll)

	router.PanicHandler = exception.PanicHandler
	return router
}
