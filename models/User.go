package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
