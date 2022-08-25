package models

import "gorm.io/gorm"

type Burger struct {
	gorm.Model

	Id          uint `gorm:"primary_key"`
	Name        string
	Brand       string
	isVegan     bool
	imageUrl    string
	Ingredients []BurgerIngredients
}