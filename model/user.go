package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" param:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Belongs  string `json:"belongs"`
	Skills   string `json:"skills"`
}

// controllerで使うメソッド--使わない予定
func (u *User) FirstById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&u)
}

func (u *User) Create() (tx *gorm.DB) {
	return DB.Create(&u)
}

func (u *User) Updates() (tx *gorm.DB) {
	return DB.Model(&u).Updates(u)
}

func (u *User) Delete() (tx *gorm.DB) {
	return DB.Delete(&u)
}

func (u *User) DeleteById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).Delete(&u)
}
