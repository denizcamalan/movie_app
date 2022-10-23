package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

type ExUser struct {
	Username	string 		`json:"username" binding:"required"`
	Password 	string 		`json:"password" binding:"required"`
}