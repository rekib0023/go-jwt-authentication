package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `gorm:"unique" json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
