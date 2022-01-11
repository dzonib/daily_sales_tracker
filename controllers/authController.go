package controllers

import (
	"github.com/dzonib/daily_sales_tracker/util"
	"strconv"
	"time"

	"github.com/dzonib/daily_sales_tracker/database"
	"github.com/dzonib/daily_sales_tracker/models"
	"github.com/gofiber/fiber/v2"
)

// var DB = database.DB

func Register(c *fiber.Ctx) error {

	var data map[string]string

	var user models.User

	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "bad request",
		})
	}

	database.DB.Where("email = ?", data["email"]).Find(&user)

	if user.Email == data["email"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email taken, please choose another",
		})
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	user = models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).Find(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
		// frontend won't be able to access
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func Logout(c *fiber.Ctx) error {
	// overwrite jwt cookie
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
		// frontend won't be able to access
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}
