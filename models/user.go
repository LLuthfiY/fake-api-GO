package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      sql.NullString
	Username  string `gorm:"uniqueIndex"`
	Password  string
	Article   []Article `gorm:"foreignKey:UserId"`
	CreatedAt *time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
