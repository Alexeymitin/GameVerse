package rtoken

import (
	"gameverse/pkg/db"
	"gameverse/pkg/model"
)

type RefreshTokenRepository interface {
	Create(token *model.RefreshToken) error
	GetByToken(token string) (*model.RefreshToken, error)
	DeleteByToken(token string) error
	DeleteAllByUserID(userID string) error
}

type refreshTokenRepo struct {
	db *db.DB
}

func NewRefreshTokenRepository(db *db.DB) RefreshTokenRepository {
	return &refreshTokenRepo{db}
}

func (r *refreshTokenRepo) Create(token *model.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *refreshTokenRepo) GetByToken(token string) (*model.RefreshToken, error) {
	var rt model.RefreshToken

	result := r.db.Where("token = ?", token).First(&rt)
	if result.Error != nil {
		return nil, result.Error
	}

	return &rt, nil
}

func (r *refreshTokenRepo) DeleteByToken(token string) error {
	return r.db.Unscoped().Where("token = ?", token).Delete(&model.RefreshToken{}).Error
}

func (r *refreshTokenRepo) DeleteAllByUserID(userID string) error {
	return r.db.Unscoped().Where("user_id = ?", userID).Delete(&model.RefreshToken{}).Error
}
