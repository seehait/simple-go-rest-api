package models

import "github.com/jinzhu/gorm"

// User represents user model.
type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Email    string `gorm:"type:varchar(191);UNIQUE_INDEX"`
	Password string `gorm:"column:password"`
}
