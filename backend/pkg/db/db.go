package db

import (
	"gameverse/pkg/configs"
	"gameverse/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDb(config *configs.Config) *DB {
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	err = db.AutoMigrate(&model.User{}, &model.RefreshToken{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return &DB{db}
}
