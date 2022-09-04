package models

import "gorm.io/gorm"

type BurgerIngredient struct {
	gorm.Model

	Id           uint `gorm:"primary_key"`
	BurgerId     uint
	IngredientId uint
	Amount       int
	Burger       Burger     `gorm:"foreignKey:BurgerId"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientId"`
}
