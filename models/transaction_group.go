package models

import (
	"gorm.io/gorm"
)

type TransactionGroup struct {
	gorm.Model
	Name        string `json:"name"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
}
