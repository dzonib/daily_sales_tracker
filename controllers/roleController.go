package controllers

import (
	"github.com/dzonib/daily_sales_tracker/database"
	"github.com/dzonib/daily_sales_tracker/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDTO models.RoleCreateDTO

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	permissions := make([]models.Permission, len(roleDTO.Permissions))

	for i, permissionId := range roleDTO.Permissions {

		permissions[i] = models.Permission{
			Id: uint(permissionId),
		}
	}

	role := models.Role{
		Name:        roleDTO.Name,
		Permissions: permissions,
	}

	database.DB.Preload("Permissions").Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var role = models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO models.RoleCreateDTO

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	permissions := make([]models.Permission, len(roleDTO.Permissions))

	for i, permissionId := range roleDTO.Permissions {

		permissions[i] = models.Permission{
			Id: uint(permissionId),
		}
	}

	var hackForDelete interface{}

	// remove previous role permissions
	// because we don't have model, we have to directly specify the table name in "Table" method
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&hackForDelete)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDTO.Name,
		Permissions: permissions,
	}

	database.DB.Model(&role).Preload("Permissions").Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var role = models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}
