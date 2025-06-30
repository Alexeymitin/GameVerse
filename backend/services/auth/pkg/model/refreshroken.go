package model

import (
	"gameverse/pkg/model"
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	model.BaseModel
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Token     string    `gorm:"type:varchar(255);not null;unique"`
	ExpiresAt time.Time
}
