package services

import (
	"github.com/josealvaradoo/expenses/models"
	"github.com/josealvaradoo/expenses/repositories"
)

func GetAllCurrencies() []models.Currency {
	var data []models.Currency
	repositories.List(&data)

	return data
}

func GetCurrencyById(id int) models.Currency {
	var data models.Currency
	repositories.Get(&data, id)

	return data
}

func CreateCurrency(data models.Currency) models.Currency {
	response := data
	repositories.Create(&response)
	return response
}

func UpdateCurrency(data models.Currency) models.Currency {
	response := data
	repositories.Update(&response)
	return response
}

func DeleteCurrency(id int) models.Currency {
	var data models.Currency
	repositories.Delete(data, id)
	return data
}
