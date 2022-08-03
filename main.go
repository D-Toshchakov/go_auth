package main

import (
	"github.com/D-Toshchakov/React-JWT-tutorial/database"
	"github.com/D-Toshchakov/React-JWT-tutorial/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
