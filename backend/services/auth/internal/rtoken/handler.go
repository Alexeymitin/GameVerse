package rtoken

import (
	"gameverse/pkg/configs"
	"gameverse/pkg/errmsg"
	"gameverse/pkg/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RefreshTokenHandler interface {
	Refresh() http.HandlerFunc
	Logout() http.HandlerFunc
	LogoutAll() http.HandlerFunc
}

type RefreshTokenHandlerDeps struct {
	RefreshTokenService RefreshTokenService
	Config              *configs.Config
}

type refreshTokenHandler struct {
	RefreshTokenService RefreshTokenService
	Config              *configs.Config
}

func NewRefreshTokenHandler(deps *RefreshTokenHandlerDeps) RefreshTokenHandler {
	return &refreshTokenHandler{
		RefreshTokenService: deps.RefreshTokenService,
		Config:              deps.Config,
	}
}

func (h *refreshTokenHandler) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("refresh_token")
		if err != nil || cookie.Value == "" {
			http.Error(w, errmsg.ErrRefreshTokenRequired, http.StatusUnauthorized)
			return
		}

		newRefreshToken, userID, err := h.RefreshTokenService.Rotate(cookie.Value)
		if err != nil {
			http.Error(w, errmsg.ErrInvalidOrExpiredToken, http.StatusUnauthorized)
			return
		}

		accessToken, err := utils.GenerateAccessToken(userID, h.Config.Jwt.AccessTTL, h.Config.Jwt.SecretKey)
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
			Value:    newRefreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(ttl),
		})

		resp := RefreshResponse{AccessToken: accessToken}

		utils.EncodeJSON(w, resp, http.StatusOK)
	}
}

func (h *refreshTokenHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		cookie, err := req.Cookie("refresh_token")
		if err == nil && cookie.Value != "" {
			_ = h.RefreshTokenService.Revoke(cookie.Value)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			MaxAge:   -1,
		})

		w.WriteHeader(http.StatusOK)
	}
}

func (h *refreshTokenHandler) LogoutAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		cookie, err := req.Cookie("refresh_token")
		if err == nil && cookie.Value != "" {
			userID, err := h.RefreshTokenService.Validate(cookie.Value)
			if err == nil && userID != uuid.Nil {
				_ = h.RefreshTokenService.RevokeAll(userID)
			}
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			MaxAge:   -1,
		})

		w.WriteHeader(http.StatusOK)
	}
}
