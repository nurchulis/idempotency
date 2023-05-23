package entity

import (
	"gorm.io/gorm"
	"time"
)

// User is a struct
type User struct {
	ID        uint      `gorm:"not null;uniqueIndex;primaryKey"`
	Name      string    `gorm:"size:100;not null;"`
	Email     string    `gorm:"size:100;not null;uniqueIndex;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

// TableName is a method
func (u User) TableName() string {
	return "users"
}
