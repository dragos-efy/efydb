package entities

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserID  uint
	ThemeID uint
	Score   int
}
