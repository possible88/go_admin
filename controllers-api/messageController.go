package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_admin/database"
	"go_admin/models"
	"go_admin/util"
	"strconv"
	"time"
)

func CreatMessage(c *fiber.Ctx) error {
	var data map[string]string

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	opened := "no"
	viewed := "no"

	database.DB.Where("id = ?", userId).Take(&user)

	message := models.Message{
		UserTo:   data["user_to"],
		UserFrom: user.UserName,
		Body:     data["body"],
		Date:     time.Now(),
		Opened:   opened,
		Viewed:   viewed,
	}

	if err := c.BodyParser(&message); err != nil {
		return err
	}

	database.DB.Create(&message)

	return c.JSON(message)
}

func GetMessage(c *fiber.Ctx) error {
	user_to := c.Params("user_to")

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Where("id = ?", userId).Take(&user)

	var message []models.Message

	database.DB.Where("user_from = ? AND user_to = ?", user.UserName, user_to).Or("user_from = ? AND user_to = ?", user_to, user.UserName).Find(&message)

	return c.JSON(fiber.Map{
		"data": message,
	})
}

//func getLatestMessage(userloggedin, user2 string) string {
//
//	message := models.Message{}
//
//	database.DB.Where("user_from = ? AND user_to = ?", userloggedin, user2).Or("user_from = ? AND user_to = ?", user2, userloggedin).Order("id desc").Find(&message)
//
//}
