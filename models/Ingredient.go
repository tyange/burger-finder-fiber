package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model

	Id          uint `gorm:"primary_key"`
	Name        string
	Kind        string
	BurgersWith []BurgerIngredient
}
