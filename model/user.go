package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" param:"id"`
	Name     string `json:"name" gorm:"type:varchar(30);not null"`
	Password string `json:"password" gorm:"not null"`
	Belongs  string `json:"belongs" gorm:"type:varchar(100);not null"`
	Skills   string `json:"skills" gorm:"type:varchar(200);not null"`
}

// controllerで使うメソッド
//使わないかも
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
