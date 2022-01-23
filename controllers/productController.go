package controllers

import (
	"github.com/dzonib/daily_sales_tracker/database"
	"github.com/dzonib/daily_sales_tracker/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllProducts(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))

	// &models.Product is entity because 2 methods we added for pagination (count and total), so we can pass it as arg.
	// its generic and it works great for user and products
	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product = models.Product{
		Id: uint(id),
	}

	// preload add associated data to query
	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product = models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product = models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return nil
}
