package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name    string
	Address Address `gorm:"embedded"`
}
