package handlers

import (
	"net/http"

	"github.com/gofiber/fiber"

	"github.com/josealvaradoo/expenses/database"
	"github.com/josealvaradoo/expenses/models"
	"github.com/josealvaradoo/expenses/services"
	"github.com/josealvaradoo/expenses/utils"
)

func GetCurrencies(c *fiber.Ctx) {
	var data []models.Currency
	currencies := services.GetAllCurrencies()

	if len(currencies) == 0 {
		data = make([]models.Currency, 0)
	} else {
		data = currencies
	}

	c.JSON(utils.Response{
		Data:    data,
		Success: true,
		Error:   "",
	})
}

func GetCurrency(c *fiber.Ctx) {
	db := database.DBConnection
	var currency models.Currency

	result := db.First(&currency)

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "Doesn't exist the currency with id: " + c.Params("id"),
		})
		return
	}

	c.JSON(utils.Response{
		Data:    currency,
		Success: true,
		Error:   "",
	})
}

func NewCurrency(c *fiber.Ctx) {
	db := database.DBConnection

	currency := new(models.Currency)

	if err := c.BodyParser(currency); err != nil {
		c.Status(http.StatusBadRequest).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "It was not possible to create a currency because some fields are missing",
		})
	}

	db.Create(&currency)

	c.Status(http.StatusCreated).JSON(utils.Response{
		Data:    currency,
		Success: true,
		Error:   "",
	})
}

func UpdateCurrency(c *fiber.Ctx) {
	db := database.DBConnection

	var currency models.Currency

	result := db.First(&currency, c.Params("id"))

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "Doesn't exist the currency with id: " + c.Params("id"),
		})
		return
	}

	if err := c.BodyParser(&currency); err != nil {
		c.Status(http.StatusBadRequest).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "It was not possible to update a currency because some fields are missing",
		})
		return
	}

	db.Updates(&currency)

	c.JSON(utils.Response{
		Data:    currency,
		Success: true,
		Error:   "",
	})
}

func DeleteCurrency(c *fiber.Ctx) {
	db := database.DBConnection
	var currency models.Currency

	result := db.First(&currency, c.Params("id"))

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound).JSON(utils.Response{
			Data:    nil,
			Success: false,
			Error:   "Doesn't exist the currency with id: " + c.Params("id"),
		})
		return
	}

	db.Delete(&currency, c.Params("id"))

	c.Status(http.StatusNoContent)
}
