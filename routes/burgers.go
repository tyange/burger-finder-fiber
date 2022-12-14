package routes

import (
	"burger-finder-fiber/database"
	"burger-finder-fiber/models"
	"burger-finder-fiber/routes/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllBurgers(c *fiber.Ctx) error {
	var burgers []models.Burger

	database.DBConn.Find(&burgers)

	return c.Status(200).JSON(burgers)
}

func FindOneBurger(c *fiber.Ctx) error {
	burger := models.Burger{}
	id := c.Params("id")

	foundBurger := database.DBConn.First(&burger, id)

	if err := foundBurger.Error; err != nil {
		return utils.ErrorHandlder(c, err)
	}

	return c.Status(200).JSON(burger)
}

func AddBurger(c *fiber.Ctx) error {
	burger := models.Burger{}

	if err := c.BodyParser(&burger); err != nil {
		return utils.ErrorHandlder(c, err)
	}

	database.DBConn.Create(&burger)

	return c.Status(200).JSON(burger)
}

func AddBurgerIngredients(c *fiber.Ctx) error {
	bgId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return utils.ErrorHandlder(c, err)
	}

	type IngredientsData struct {
		Ingredients []struct {
			Amount       int
			IngredientId uint
		}
	}

	var ingredientsData IngredientsData

	if err := c.BodyParser(&ingredientsData); err != nil {
		return utils.ErrorHandlder(c, err)
	}

	for _, v := range ingredientsData.Ingredients {
		burgerIngredient := models.BurgerIngredient{Amount: v.Amount, BurgerId: uint(bgId), IngredientId: v.IngredientId}
		database.DBConn.Create(&burgerIngredient)
	}

	return c.Status(200).JSON(ingredientsData)
}

func UpdateBurger(c *fiber.Ctx) error {
	burger := models.Burger{}
	updatedData := models.Burger{}
	id := c.Params("id")

	if err := c.BodyParser(&updatedData); err != nil {
		return utils.ErrorHandlder(c, err)
	}

	database.DBConn.First(&burger, id).Updates(&updatedData)

	return c.Status(200).JSON(burger)
}

func DeleteBurger(c *fiber.Ctx) error {
	burger := models.Burger{}
	id := c.Params("id")
	targetBurger := database.DBConn.First(&burger, id)

	if err := targetBurger.Error; err != nil {
		return utils.ErrorHandlder(c, err)
	}

	targetBurger.Delete(&burger)

	return c.Status(200).JSON("deleted")
}
