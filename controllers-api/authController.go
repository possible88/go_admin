package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_admin/database"
	"go_admin/models"
	"go_admin/util"
	"math/rand"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	token := rand.ExpFloat64()

	user := models.User{
		FirstName:  data["first_name"],
		LastName:   data["last_name"],
		UserName:   data["user_name"],
		Email:      data["email"],
		Country:    data["country"],
		ProfilePic: data["profile_pic"],
		Phone:      data["phone"],
		Date:       time.Now(),
		Token:      token,
	}

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id > 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email already Exist",
		})
	}
	database.DB.Where("user_name = ?", data["user_name"]).First(&user)

	if user.Id > 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Username already Exist",
		})
	}
	database.DB.Where("phone = ?", data["phone"]).First(&user)
	if user.Id > 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Phone number already Exist",
		})
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "Registration successful",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
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

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id:        uint(userId),
		Skill: data["skill"],
		AboutMe: data["about_me"],
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(fiber.Map{
		"message": "Update Complete",
	})

}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)

}

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}

		url := "http://localhost:8000/api/uploads/" +
			filename

		userId, _ := strconv.Atoi(id)

		user := models.User{
			Id:         uint(userId),
			ProfilePic: url,
		}

		database.DB.Model(&user).Updates(user)

	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + filename,
	})
}
