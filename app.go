package main

import (
	"log"
	_ "log"

	"burger-finder-fiber/database"
	"burger-finder-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", routes.AllIngredients)
	app.Post("/", routes.AddIngredient)
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
