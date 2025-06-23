package auth

import (
	"gameverse/pkg/configs"
	"gameverse/pkg/errmsg"
	"gameverse/pkg/utils"
	"gameverse/services/auth/internal/rtoken"
	"net/http"
	"time"
)

type AuthHandler interface {
	Login() http.HandlerFunc
	Register() http.HandlerFunc
}

type AuthHandlerDeps struct {
	AuthService         AuthService
	RefreshTokenService rtoken.RefreshTokenService
	Config              *configs.Config
}

type authHandler struct {
	AuthService         AuthService
	RefreshTokenService rtoken.RefreshTokenService
	Config              *configs.Config
}

func NewAuthHandler(deps *AuthHandlerDeps) AuthHandler {
	return &authHandler{
		AuthService:         deps.AuthService,
		Config:              deps.Config,
		RefreshTokenService: deps.RefreshTokenService,
	}
}

func (h *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := utils.DecodeJSON[RegisterRequest](req)
		if err != nil {
			http.Error(w, errmsg.ErrInvalidRequest, http.StatusBadRequest)
			return
		}

		err = utils.ValidateStruct(body)
		if err != nil {
			http.Error(w, errmsg.ErrValidationFailed, http.StatusBadRequest)
			return
		}

		user, err := h.AuthService.Register(&body)
		if err != nil {
			http.Error(w, errmsg.ErrInternalServer, http.StatusInternalServerError)
			return
		}

		accessToken, err := utils.GenerateAccessToken(user.ID, h.Config.Jwt.AccessTTL, h.Config.Jwt.SecretKey)
		if err != nil {
			http.Error(w, errmsg.ErrFailedToGenerateToken, http.StatusInternalServerError)
			return
		}

		refreshToken, err := h.RefreshTokenService.Generate(user.ID)
		if err != nil {
			http.Error(w, errmsg.ErrFailedToGenerateToken, http.StatusInternalServerError)
			return
		}

		ttl, err := time.ParseDuration(h.Config.Jwt.RefreshTTL)
		if err != nil {
			http.Error(w, errmsg.ErrFailedToParseTokenTTL, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(ttl),
		})

		resp := RegisterResponse{AccessToken: accessToken, Message: "registration successful"}

		utils.EncodeJSON(w, resp, http.StatusOK)
	}
}

func (h *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// TODO: добавить удаление рефреш токенов
		body, err := utils.DecodeJSON[LoginRequest](req)
		if err != nil {
			http.Error(w, errmsg.ErrInvalidRequest, http.StatusBadRequest)
			return
		}

		err = utils.ValidateStruct(body)
		if err != nil {
			http.Error(w, errmsg.ErrValidationFailed, http.StatusBadRequest)
			return
		}

		accessToken, userID, err := h.AuthService.Login(&body)
		if err != nil {
			http.Error(w, errmsg.ErrInternalServer, http.StatusInternalServerError)
			return
		}

		refreshToken, err := h.RefreshTokenService.Generate(userID)
		if err != nil {
			http.Error(w, errmsg.ErrFailedToGenerateToken, http.StatusInternalServerError)
			return
		}

		ttl, err := time.ParseDuration(h.Config.Jwt.RefreshTTL)
		if err != nil {
			http.Error(w, errmsg.ErrFailedToParseTokenTTL, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(ttl),
		})

		resp := LoginResponse{AccessToken: accessToken}

		utils.EncodeJSON(w, resp, http.StatusOK)
	}
}
