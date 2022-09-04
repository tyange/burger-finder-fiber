package main

import (
	"log"

	"burger-finder-fiber/database"
	"burger-finder-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/ingredients", routes.AllIngredients)
	app.Get("/ingredients/:id", routes.FindOneIngredient)
	app.Post("/ingredients", routes.AddIngredient)
	app.Put("/ingredients/:id", routes.UpdateIngredient)
	app.Delete("/ingredients/:id", routes.DeleteIngredient)
	app.Get("/burgers", routes.AllBurgers)
	app.Get("/burgers/:id", routes.FindOneBurger)
	app.Post("/burgers", routes.AddBurger)
	app.Post("/burgers/:id", routes.AddBurgerIngredients)
	app.Put("/burgers/:id", routes.UpdateBurger)
	app.Delete("/burgers/:id", routes.DeleteBurger)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
