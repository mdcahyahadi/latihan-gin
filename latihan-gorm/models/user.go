package models

import (
	"time"
)

// User represents the model of an user
type User struct {
	// gorm.Model
	ID        uint   `gorm:"Primary Key"`
	Email     string `gorm:"unique;not null;type:varchar(191)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
