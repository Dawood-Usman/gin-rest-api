package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null;default:null"`
	Author string `json:"author" gorm:"not null;default:null"`
}
