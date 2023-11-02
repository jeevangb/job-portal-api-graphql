package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Company Company `gorm:"foreignKey:Cid"`
	Cid     uint    `json:"cid"`
	Title   string  `json:"title"`
	Desc    string  `json:"desc"`
}
