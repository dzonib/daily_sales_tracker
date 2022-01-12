package main

import (
	"github.com/dzonib/daily_sales_tracker/database"
	"github.com/dzonib/daily_sales_tracker/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		// for FE cookies
		AllowCredentials: true,
	}))

	routes.Setup(app)
	// gin -p 8080
	err := app.Listen(":3001")

	if err != nil {
		return
	}
}
