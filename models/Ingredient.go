package models

import (
	"gorm.io/gorm"
	"time"
)

type Ingredient struct {
	gorm.Model

	Id          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Kind        string
	BurgersWith []BurgerIngredient
}
