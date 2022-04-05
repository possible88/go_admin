package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_admin/database"
	"go_admin/models"
	"math"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5

	offset := (page - 1) * limit
	var total int64

	var users []models.User

	database.DB.Offset(offset).Limit(limit).Find(&users)

	database.DB.Model(&models.User{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": users,
		"mata": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func GetUser(c *fiber.Ctx) error {
	username := c.Params("user_name")

	var user models.User

	database.DB.Where("user_name = ?", username).Find(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
