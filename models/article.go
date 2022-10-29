package models

import (
	"database/sql"
	"time"
)

type Article struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Content   string
	UserId    string
	User      *User `gorm:"foreignKey:UserId;reference:ID"`
	CreatedAt *time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
