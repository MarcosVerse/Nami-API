package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey;column:id"     json:"id"`
	Name     string `gorm:"column:name"              json:"name"`
	Email    string `gorm:"column:email;uniqueIndex" json:"email"`
	Password string `gorm:"column:password"          json:"-"`

	CreatedAt time.Time      `gorm:"column:created_at"       json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"       json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (User) TableName() string {
	return "users"
}
