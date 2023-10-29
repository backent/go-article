package main

import (
	"net/http"

	controllersUser "github.com/backent/go-article/controllers/user"
	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/libs"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	servicesUser "github.com/backent/go-article/services/user"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := libs.Initiate()
	validate := validator.New(validator.WithRequiredStructEnabled())

	userRepository := repositoriesUser.NewRepositoryMysqlImpl()
	userService := servicesUser.NewServiceUser(db, userRepository, validate)
	userController := controllersUser.NewControllerUser(userService)

	router := httprouter.New()

	router.POST("/", userController.Create)
	router.PUT("/:id", userController.Update)
	router.DELETE("/:id", userController.Delete)
	router.GET("/:id", userController.FindById)
	router.GET("/", userController.FindAll)

	router.PanicHandler = exception.PanicHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
