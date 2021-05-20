package models

import (
	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	Name   string  `json:"name"`
	Value  float64 `json:"value"`
	Factor float64 `json:"factor"`
	Icon   string  `json:"icon"`
}
