package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_admin/database"
	"go_admin/models"
	"go_admin/util"
	"math"
	"strconv"
	"time"
)

func AllJob(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5

	offset := (page - 1) * limit
	var total int64

	var jobs []models.Job

	database.DB.Order("id desc").Offset(offset).Limit(limit).Find(&jobs)

	database.DB.Model(&models.Job{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": jobs,
		"mata": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateJob(c *fiber.Ctx) error {
	var data map[string]string

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Where("id = ?", userId).Take(&user)

	job := models.Job{
		Title:       data["title"],
		Description: data["description"],
		Company:     data["company"],
		Website:     data["website"],
		Period:      data["period"],
		JobNature:   data["jobnature"],
		Skill:       data["skill"],
		Education:   data["education"],
		AddedBy:     user.UserName,
		State:       data["state"],
		Country:     data["country"],
		Payment:       data["payment"],
		Date:        time.Now(),
	}

	if err := c.BodyParser(&job); err != nil {
		return err
	}

	database.DB.Create(&job)

	return c.JSON(fiber.Map{
		"message": "Created Successful",
	})
}

func GetJob(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	job := models.Job{
		Id: uint(id),
	}

	database.DB.Find(&job)

	return c.JSON(job)
}

func UpdateJob(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	job := models.Job{
		Id: uint(id),
	}

	if err := c.BodyParser(&job); err != nil {
		return err
	}

	database.DB.Model(&job).Updates(job)

	return c.JSON(job)
}

func DeleteJob(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	job := models.Job{
		Id: uint(id),
	}

	database.DB.Delete(&job)

	return nil
}
