package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Type               string  `json:"type"`
	CurrentBalance     float64 `json:"current_balance"`
	Value              float64 `json:"value"`
	CurrencyID         int     `json:"currency_id"`
	WalletID           int     `json:"wallet_id"`
	Description        string  `json:"description"`
	SourceID           int     `json:"source_id"`
	TransactionGroupID int     `json:"transaction_group_id"`
}
