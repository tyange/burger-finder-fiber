package models

import "gorm.io/gorm"

type BurgerIngredients struct {
	gorm.Model

	Id             uint `gorm:"primary_key"`
	BurgerId       int
	IngredientId   int
	IngredientName string
	Amount         int
	Burger         Burger
	Ingredient     Ingredient
}
