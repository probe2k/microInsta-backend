package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID	uint	`gorm:"primary_key,autoIncrement"`
	Name string	`gorm:"column:name"`
	Mobile uint `gorm:"unique,column:mobile"`
	Address string `gorm:"column:address"`
	postCount uint `gorm:"column:count"`
}
