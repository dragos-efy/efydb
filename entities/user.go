package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `validate:"required,min=3,max=32" json:"name"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
	// 0 = User, 1 = Admin, 2 = Owner
	Role  uint   `json:"role"`
	Token string `json:"token"`
}
