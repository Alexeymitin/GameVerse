package main

import (
	"gameverse/configs"
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

	println("Backend service is running...")

	server.ListenAndServe()
}
