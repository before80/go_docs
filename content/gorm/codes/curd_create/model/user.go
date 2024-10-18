package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint       `gorm:"primaryKey,autoIncrement,not null" json:"id"`
	Name         string     `json:"name"`
	Email        *string    `json:"email"`
	Age          uint8      `json:"age"`
	Birthday     *time.Time `json:"birthday"`
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	gorm.Model
}
