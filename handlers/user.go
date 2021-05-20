package handlers

import (
	"net/http"

	"github.com/gofiber/fiber"

	"github.com/josealvaradoo/expenses/database"
	"github.com/josealvaradoo/expenses/models"
	"github.com/josealvaradoo/expenses/utils"
)

func GetUsers(c *fiber.Ctx) {
	var data []models.User
	var users []models.User

	db := database.DBConnection
	db.Find(&users)

	if len(users) == 0 {
		data = make([]models.User, 0)
	} else {
		data = users
	}

	c.JSON(utils.Response{
		Data:    data,
		Success: true,
		Error:   "",
	})
}

func GetUser(c *fiber.Ctx) {
	db := database.DBConnection
	var user models.User

	result := db.First(&user, c.Params("id"))

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "Doesn't exist the user with id: " + c.Params("id"),
		})
		return
	}

	c.JSON(utils.Response{
		Data:    user,
		Success: true,
		Error:   "",
	})
}

func NewUser(c *fiber.Ctx) {
	db := database.DBConnection

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(http.StatusBadRequest).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "It was not possible to create an user because some fields are missing",
		})
	}

	db.Create(&user)

	c.Status(http.StatusCreated).JSON(utils.Response{
		Data:    user,
		Success: true,
		Error:   "",
	})
}

func UpdateUser(c *fiber.Ctx) {
	db := database.DBConnection

	var user models.User

	result := db.First(&user, c.Params("id"))

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "Doesn't exist the user with id: " + c.Params("id"),
		})
		return
	}

	if err := c.BodyParser(&user); err != nil {
		c.Status(http.StatusBadRequest).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "It was not possible to update an user because some fields are missing",
		})
		return
	}

	db.Updates(&user)

	c.JSON(utils.Response{
		Data:    user,
		Success: true,
		Error:   "",
	})
}

func DeleteUser(c *fiber.Ctx) {
	db := database.DBConnection
	var user models.User

	result := db.First(&user, c.Params("id"))

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "Doesn't exist the user with id: " + c.Params("id"),
		})
		return
	}

	db.Delete(&user, c.Params("id"))

	c.Status(http.StatusNoContent)
}
