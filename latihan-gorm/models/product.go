package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uint   `gorm "Primary Key"`
	Name      string `gorm: "not null; type:varchar(191)"`
	BRand     string `gorm: "not null; type:varchar(191)"`
	UserID    uint   `gorm: "foreignkey:User`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create product")
	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}
	return
}
