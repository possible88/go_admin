package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_admin/database"
	"go_admin/models"
	"go_admin/util"
	"strconv"
	"time"
)

func CreatComment(c *fiber.Ctx) error {
	var data map[string]string

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Where("id = ?", userId).Take(&user)


	comment := models.Comment{
		PostedTo:   data["posted_to"],
		PostedBy: user.UserName,
		Body:     data["body"],
		Date:     time.Now(),
		FirstName: user.FirstName,
		LastName: user.LastName,
		ProfilePic: user.ProfilePic,
		UserName: user.UserName,
	}

	if err := c.BodyParser(&comment); err != nil {
		return err
	}


	database.DB.Create(&comment)

	
	return c.JSON(comment)
}


func GetComment(c *fiber.Ctx) error {
	posted_to := c.Params("posted_to")

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Where("id = ?", userId).Take(&user)

	var comment []models.Comment

	database.DB.Where("posted_to = ?",  posted_to).Find(&comment)

	return c.JSON(fiber.Map{
		"data": comment,
	})
}