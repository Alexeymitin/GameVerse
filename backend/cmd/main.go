package main

import (
	"gameverse/pkg/configs"
	"gameverse/pkg/db"
	"net/http"
)

func app() http.Handler {
	config := configs.LoadConfig()
	db.NewDb(config)
	router := http.NewServeMux()

	return router
}

func main() {
	app := app()
	server := http.Server{
		Addr:    ":50059",
		Handler: app,
	}

	println("Auth service is running on port:50059")

	server.ListenAndServe()
}
