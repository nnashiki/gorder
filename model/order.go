package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID uint `json:"id"`
	UserID uint
	User User
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (o *Order) FirstById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id ).First(&o)
}

func (o *Order) Create() (tx *gorm.DB) {
	return DB.Create(&o)
}