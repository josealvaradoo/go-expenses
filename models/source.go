package models

import (
	"gorm.io/gorm"
)

type Source struct {
	gorm.Model
	Name string `json:"name"`
	Icon string `json:"icon"`
}
