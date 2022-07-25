package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NickName string `gorm:"not null" json:"nick_name"`
	Email    string `gorm:"not null; unique_index" json:"email"`
	Tasks    []Task `json:"tasks"`
}
