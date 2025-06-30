package rtoken

import (
	"errors"
	"gameverse/pkg/configs"
	"gameverse/services/auth/pkg/model"

	"gameverse/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type RefreshTokenService interface {
	Generate(userID uuid.UUID) (string, error)
	Validate(token string) (uuid.UUID, error)
	Revoke(token string) error
	RevokeAll(userID uuid.UUID) error
	Rotate(token string) (newToken string, userID uuid.UUID, err error)
}

type RefreshTokenDeps struct {
	Repo   RefreshTokenRepository
	Config *configs.Config
}

type refreshTokenService struct {
	repo   RefreshTokenRepository
	config *configs.Config
}

func NewRefreshTokenService(deps *RefreshTokenDeps) RefreshTokenService {
	return &refreshTokenService{
		repo:   deps.Repo,
		config: deps.Config,
	}
}

func (r *refreshTokenService) Generate(userID uuid.UUID) (string, error) {
	return r.createToken(userID)
}

func (r *refreshTokenService) Validate(token string) (uuid.UUID, error) {
	rt, err := r.repo.GetByToken(token)
	if err != nil {
		return uuid.Nil, err
	}

	if rt.ExpiresAt.Before(time.Now()) {
		_ = r.repo.DeleteByToken(token)
		return uuid.Nil, errors.New("token expired")
	}

	return rt.UserID, nil
}

func (s *refreshTokenService) Revoke(token string) error {
	return s.repo.DeleteByToken(token)
}

func (s *refreshTokenService) RevokeAll(userID uuid.UUID) error {
	return s.repo.DeleteAllByUserID(userID.String())
}

func (r *refreshTokenService) Rotate(token string) (newToken string, userID uuid.UUID, err error) {
	rt, err := r.repo.GetByToken(token)
	if err != nil {
		return "", uuid.Nil, err
	}

	if rt.ExpiresAt.Before(time.Now()) {
		_ = r.repo.DeleteByToken(token)
		return "", uuid.Nil, errors.New("token expired")
	}

	// Удаляем старый токен (ротация)
	_ = r.repo.DeleteByToken(token)

	// Генерируем новый токен
	newToken, err = r.createToken(rt.UserID)
	if err != nil {
		return "", uuid.Nil, err
	}

	return newToken, rt.UserID, nil
}

func (r *refreshTokenService) createToken(userID uuid.UUID) (string, error) {
	token, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		return "", err
	}

	ttl, err := time.ParseDuration(r.config.Jwt.RefreshTTL)
	if err != nil {
		panic("invalid JWT_REFRESH_TTL format in config: " + err.Error())
	}

	rt := &model.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(ttl),
	}
	if err := r.repo.Create(rt); err != nil {
		return "", err
	}

	return token, nil
}
