package controllers

import (
	"github.com/dzonib/daily_sales_tracker/database"
	"github.com/dzonib/daily_sales_tracker/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))

	//limit := 3

	//offset := (page - 1) * limit

	//var total int64
	//
	//var users []models.User
	//
	//database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	//
	//database.DB.Model(&models.User{}).Count(&total)

	//return c.JSON(fiber.Map{
	//	"data": users,
	//	"meta": fiber.Map{
	//		"total":     total,
	//		"page":      page,
	//		"last_page": math.Ceil(float64(int(total) / limit)),
	//	},
	//})

	// &models.User is entity because 2 methods we added for pagination (count and total), so we can pass it as arg.
	// its generic and it works great for user and products
	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var user = models.User{
		Id: uint(id),
	}

	// preload add associated data to query
	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var user = models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var user = models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
