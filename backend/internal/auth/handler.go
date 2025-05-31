package auth

import (
	"fmt"
	"gameverse/configs"
	"gameverse/pkg/request"
	"gameverse/pkg/response"
	"net/http"
)

type AuthHandler struct {
	Config *configs.Config
}

type AuthHandlerDeps struct {
	Config *configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("Login endpoint hit with secret key: %s\n", h.Config.Auth.SecretKey)

		body, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			return
		}

		fmt.Println("Login request body:", body)

		data := LoginResponse{
			Token: "123",
		}

		response.Json(w, data, http.StatusOK)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Print("Register endpoint hit\n")

		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			return
		}

		fmt.Println("Login request body:", body)

		data := LoginResponse{
			Token: "123",
		}

		response.Json(w, data, http.StatusOK)
	}
}
