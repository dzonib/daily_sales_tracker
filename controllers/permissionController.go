package controllers

import (
	"github.com/dzonib/daily_sales_tracker/database"
	"github.com/dzonib/daily_sales_tracker/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	// alt + j select next occurrence
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
