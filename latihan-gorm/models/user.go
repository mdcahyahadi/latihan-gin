package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm: "Primary Key"`
	Email     string `gorm: "unique;not null;type:varchar(191)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
