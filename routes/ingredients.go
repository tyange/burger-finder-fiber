package routes

import (
	"burger-finder-fiber/database"
	"burger-finder-fiber/models"
	"burger-finder-fiber/routes/utils"
	"github.com/gofiber/fiber/v2"
)

func AllIngredients(c *fiber.Ctx) error {
	var ingredients []models.Ingredient

	database.DBConn.Find(&ingredients)

	return c.Status(200).JSON(ingredients)
}

func FindOneIngredient(c *fiber.Ctx) error {
	ingredient := models.Ingredient{}
	id := c.Params("id")

	foundIngredient := database.DBConn.First(&ingredient, id)

	if err := foundIngredient.Error; err != nil {
		return utils.ErrorHandlder(c, err)
	}

	return c.Status(200).JSON(ingredient)
}

func AddIngredient(c *fiber.Ctx) error {
	ingredient := models.Ingredient{}

	if err := c.BodyParser(&ingredient); err != nil {
		return utils.ErrorHandlder(c, err)
	}

	database.DBConn.Create(&ingredient)

	return c.Status(200).JSON(ingredient)
}

func UpdateIngredient(c *fiber.Ctx) error {
	ingredient := models.Ingredient{}
	updatedData := new(models.Ingredient)
	id := c.Params("id")

	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	targetIngredient := database.DBConn.First(&ingredient, id)

	if err := targetIngredient.Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	targetIngredient.Updates(updatedData)

	return c.Status(200).JSON(ingredient)
}

func DeleteIngredient(c *fiber.Ctx) error {
	ingredient := models.Ingredient{}
	id := c.Params("id")
	targetIngredient := database.DBConn.First(&ingredient, id)

	if err := targetIngredient.Error; err != nil {
		return utils.ErrorHandlder(c, err)
	}

	targetIngredient.Delete(&ingredient)

	return c.Status(200).JSON("deleted")
}
