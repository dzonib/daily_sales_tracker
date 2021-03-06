package routes

import (
	"github.com/dzonib/daily_sales_tracker/controllers"
	"github.com/dzonib/daily_sales_tracker/util/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// public
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// middleware
	app.Use(middleware.IsAuthenticated)

	//	authenticated only
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)
}
