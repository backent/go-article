package main

import (
	"net/http"

	"github.com/backent/go-article/injector"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	router := injector.InitializeRouter()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
