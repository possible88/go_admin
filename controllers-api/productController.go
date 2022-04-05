package controllers

import (
	// "context"
	// "encoding/json"
	"context"
	"encoding/json"
	"fmt"
	"go_admin/database"
	"go_admin/models"
	"go_admin/util"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	// "strings"
	"time"
)

func AllProduct(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 30

	offset := (page - 1) * limit
	var total int64

	var products []models.Product

	database.DB.Order("id desc").Offset(offset).Limit(limit).Find(&products)

	database.DB.Model(&models.Product{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": products,
		"mata": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var data map[string]string

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Where("id = ?", userId).Take(&user)

	rand := rand.ExpFloat64()

	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image[]"]
	
	var filename string 


	for _, file := range files {

          filename = "http://localhost:8000/api/posts/" +file.Filename

		 fmt.Println(filename)

		if err := c.SaveFile(file, "./posts/"+file.Filename); err != nil {
			return err
		}

		productimg := models.Productimg{
			Image:  filename,
			ImgRef: rand,
		}
	
		database.DB.Create(&productimg)

	}

	// str := strings.Join(filename," ")
    // fmt.Println(str)

	// url := "http://localhost:8000/api/posts/" +
	// 	filename

	

	product := models.Product{
		Category:      data["category"],
		Title:         data["tile"],
		Description:   data["description"],
		Itemcondition: data["itemcondition"],
		AddedBy:       user.UserName,
		State:         data["state"],
		Country:       data["country"],
		Image:        filename,
		Price:       data["price"],
		Date:          time.Now(),
		ImgRef:        rand,
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(fiber.Map{
		"message": "Created Successful",
	})
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&product)

	return c.JSON(product)
}


func Getphoto(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Where("id = ?", id).Take(&product)

	var image []models.Productimg

	database.DB.Where("img_ref = ?", product.ImgRef).Find(&image)

	return c.JSON(fiber.Map{
		"data": image,
	})
}


func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(fiber.Map{
		"message": "Update Successful",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return nil
}

func SearchProduct(c *fiber.Ctx) error {

	var products []models.Product
	var ctx = context.Background()
	expiredTime := 30 * time.Minute

	result, err := database.Cache.Get(ctx, "products_backend").Result()

	if err != nil {
		database.DB.Order("id desc").Find(&products)

		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}

		database.Cache.Set(ctx, "product_backend", bytes, expiredTime).Err()
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	var searchproducts []models.Product

	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), lower) || strings.Contains(strings.ToLower(product.Description), lower) {
				searchproducts = append(searchproducts, product)
			}
		}
	} else {
		searchproducts = products
	}

	var total = len(searchproducts)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 30

	var data []models.Product

	if total <= page*perPage && total >= (page-1)*perPage {
		data = searchproducts[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = searchproducts[(page-1)*perPage : page*perPage]
	} else {
		data = []models.Product{}
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}

func GetProductAddedBy(c *fiber.Ctx) error {
	addedBy := c.Params("added_by")

	var product []models.Product

	database.DB.Where("added_by = ?", addedBy).Order("id desc").Find(&product)

	var total = len(product)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 30

	var data []models.Product

	if total <= page*perPage && total >= (page-1)*perPage {
		data = product[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = product[(page-1)*perPage : page*perPage]
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}

func GetProductCategory(c *fiber.Ctx) error {
	category := c.Params("category")

	var product []models.Product

	database.DB.Where("category = ?", category).Order("id desc").Find(&product)

	var total = len(product)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 30

	var data []models.Product

	if total <= page*perPage && total >= (page-1)*perPage {
		data = product[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = product[(page-1)*perPage : page*perPage]
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}
