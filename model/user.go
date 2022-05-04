package model

import (
	"gorm.io/gorm"
	"go-member/data"
)

// controllerで使う関数

// structをimportした時のポインタの書き方が分からない
func (u *data.User) FirstById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&u)
}

func (u *data.User) Create() (tx *gorm.DB) {
	return DB.Create(&u)
}

func (u *data.User) Save() (tx *gorm.DB) {
	return DB.Save(&u)
}

func (u *data.User) Updates() (tx *gorm.DB) {
	return DB.Model(&u).Updates(u)
}

func (u *data.User) Delete() (tx *gorm.DB) {
	return DB.Delete(&u)
}

func (u *data.User) DeleteById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).Delete(&u)
}
