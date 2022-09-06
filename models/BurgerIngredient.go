package models

import (
	"gorm.io/gorm"
	"time"
)

type BurgerIngredient struct {
	gorm.Model

	Id           uint `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	BurgerId     uint
	IngredientId uint
	Amount       int
	Burger       Burger     `gorm:"foreignKey:BurgerId"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientId"`
}
