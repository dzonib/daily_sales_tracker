package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := 3

	offset := (page - 1) * limit

	// rewritten stuff with interfaces, much cooler, less code and more readable
	data := entity.Take(db, limit, offset)

	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}
}
