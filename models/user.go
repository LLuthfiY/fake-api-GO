package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        string `gorm:"uniqueIndex"`
	Name      sql.NullString
	Username  string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
