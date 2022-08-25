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

func FindOneIngredient(c *fiber.Ctx) error {
	var ingredient []models.Ingredient
	id := c.Params("id")

	database.DBConn.Model(&ingredient).Where("id = ?", &id)

	return c.Status(200).JSON(&ingredient)
}

func AddIngredient(c *fiber.Ctx) error {
	ingredient := new(models.Ingredient)
	if err := c.BodyParser(ingredient); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&ingredient)

	return c.Status(200).JSON(ingredient)
}

func UpdateIngredient(c *fiber.Ctx) error {
	var ingredient []models.Ingredient
	updatedData := new(models.Ingredient)
	id := c.Params("id")

	targetIngredient := database.DBConn.Model(&ingredient).Where("id = ?", id)

	if err := c.BodyParser(updatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	targetIngredient.Updates(models.Ingredient{Name: updatedData.Name, Kind: updatedData.Kind})

	return c.Status(200).JSON(updatedData)
}

func DeleteIngredient(c *fiber.Ctx) error {
	var ingredient []models.Ingredient
	id := c.Params("id")
	targetIngredient := database.DBConn.Model(&ingredient).Where("id = ?", id)

	targetIngredient.Delete(&ingredient)

	return c.Status(200).JSON("deleted")
}
