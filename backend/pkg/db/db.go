package db

import (
	"gameverse/pkg/configs"

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

	return &DB{db}
}

func (db *DB) Migrate(models ...any) {
	err := db.AutoMigrate(models...)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
