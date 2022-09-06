package models

import (
	"gorm.io/gorm"
	"time"
)

type Burger struct {
	gorm.Model

	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Brand     string
	IsVegan   bool
	ImageUrl  string
}
