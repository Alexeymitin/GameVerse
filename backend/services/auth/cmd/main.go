package main

import (
	"gameverse/pkg/configs"
	"gameverse/pkg/db"
	"gameverse/services/auth/internal/auth"
	"gameverse/services/auth/internal/rtoken"
	"gameverse/services/auth/pkg/model"

	"net/http"
)

func app() (http.Handler, *configs.Config) {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	db.Migrate(&model.User{}, model.RefreshToken{})
	router := http.NewServeMux()

	// repositories
	userRepo := auth.NewUserRepository(db)
	refreshTokenRepo := rtoken.NewRefreshTokenRepository(db)

	// services
	refreshTokenService := rtoken.NewRefreshTokenService(&rtoken.RefreshTokenDeps{
		Repo:   refreshTokenRepo,
		Config: config,
	})
	authService := auth.NewAuthService(&auth.AuthServiceDeps{
		UserRepository: userRepo,
		Config:         config,
	})

	// handlers
	refreshTokenHandler := rtoken.NewRefreshTokenHandler(&rtoken.RefreshTokenHandlerDeps{
		RefreshTokenService: refreshTokenService,
		Config:              config,
	})

	authHandler := auth.NewAuthHandler(&auth.AuthHandlerDeps{
		AuthService:         authService,
		RefreshTokenService: refreshTokenService,
		Config:              config,
	})

	// routers
	router.HandleFunc("POST /auth/refresh", refreshTokenHandler.Refresh())
	router.HandleFunc("POST /auth/logout", refreshTokenHandler.Logout())
	router.HandleFunc("POST /auth/logoutAll", refreshTokenHandler.LogoutAll())

	router.HandleFunc("POST /auth/register", authHandler.Register())
	router.HandleFunc("POST /auth/login", authHandler.Login())

	return router, config
}

func main() {
	app, config := app()
	server := http.Server{
		Addr:    ":8000",
		Handler: app,
	}

	println("Auth service is running on port:8000")

	err := server.ListenAndServeTLS(config.Ssl.SSLCertPath, config.Ssl.SSLKeyPath)
	if err != nil {
		panic(err)
	}
}
