package models

import (
	"FGA_Hacktiv8/jwt/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Fullname string    `json:"full_name" form:"full_name" gorm:"not null" valid:"required~Your full name is required"`
	Email    string    `json:"email" form:"email" gorm:"not null;uniqueIndex" valid:"required~Your email is required,email~Invalid email format"`
	Password string    `json:"password" form:"password" gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(d *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password = helpers.HashPassword(u.Password)

	return nil
}
