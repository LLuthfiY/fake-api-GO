package config

import (
	"github.com/LLuthfiY/fake-api-GO/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic("can't open sqlite")
	}
	return db
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("cannot close sqlite")
	}
	dbSQL.Close()
}
