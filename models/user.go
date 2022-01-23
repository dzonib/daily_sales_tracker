package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	// don't show in json
	Password []byte `json:"-"`
	RoleId   uint   `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (u *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	u.Password = hashedPassword
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}

func (u *User) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&User{}).Count(&total)

	return total
}

func (u *User) Take(db *gorm.DB, limit int, offset int) interface{} {

	var user []User

	db.Preload("Role").Offset(offset).Limit(limit).Find(&user)

	return user
}
