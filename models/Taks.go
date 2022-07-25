package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Tittle      string `gorm:"not null; unique_index" json:"tittle"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      uint   `json:"user_id"`
}
