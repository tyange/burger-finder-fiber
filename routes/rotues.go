package routes

import (
	"burger-finder-fiber/database"
	"burger-finder-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func AllIngredients(c *fiber.Ctx) error {
	var ingredients []models.Ingredient

	database.DBConn.Find(&ingredients)

	return c.Status(200).JSON(ingredients)
}

func AddIngredient(c *fiber.Ctx) error {
	ingredient := new(models.Ingredient)
	if err := c.BodyParser(ingredient); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&ingredient)

	return c.Status(200).JSON(ingredient)
}

func DeleteIngredient(c *fiber.Ctx) error {
	var ingredient []models.Ingredient
	id := c.Params("id")
	database.DBConn.Where("id = ?", id).Delete(&ingredient)

	return c.Status(200).JSON("deleted")
}
