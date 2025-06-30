package model

import "gameverse/pkg/model"

type User struct {
	model.BaseModel
	Email       string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password    string `gorm:"type:varchar(255);not null" json:"-"`
	FirstName   string `gorm:"type:varchar(100)" json:"first_name"`
	LastName    string `gorm:"type:varchar(100)" json:"last_name"`
	MobilePhone string `gorm:"type:varchar(20);index" json:"mobile_phone"`
}
