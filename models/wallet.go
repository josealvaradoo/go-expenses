package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Name       string  `json:"name"`
	UserID     int     `json:"user_id"`
	CurrencyID int     `json:"currency_id"`
	Balance    float64 `json:"balance"`
	Icon       float64 `json:"icon"`
}
