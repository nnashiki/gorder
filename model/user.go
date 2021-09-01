package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID		  uint `json:"id"`
	Name 	  string `json:"name" gorm:"type:varchar(255);not null"`
	Email	  string `json:"email" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) FirstById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&u)
}

func (u *User) Create() (tx *gorm.DB) {
	return DB.Create(&u)
}

func (u *User) Updates(newu User) (tx *gorm.DB) {
	return DB.Model(&u).Updates(newu)
}
