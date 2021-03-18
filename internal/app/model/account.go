package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Phone string `gorm:"type:varchar(15);unique;not null;comment:手机号"`
	Email string `gorm:"unique;comment:邮箱"`
	Username string `gorm:"unique";comment:用户名"`
	Password string `gorm:"comment:密码"`
	Status int `gorm:"type:tinyint(1);default:0;comment:账户状态"`
}
