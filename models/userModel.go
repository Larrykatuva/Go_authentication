package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string
	Email       string
	PhoneNumber string
	Password    string
	Inactive    bool
	Verified    bool
}
