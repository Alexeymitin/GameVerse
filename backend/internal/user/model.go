package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"index"`
	Name     string
	Password string
}

// func NewUser(email string) *User {
// 	user := &User{
// 		Email: email,
// 	}

// 	return user
// }
