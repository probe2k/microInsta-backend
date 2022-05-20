package model

import (
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	PID uint `gorm:"primary_key,autoIncrement"`
	ID uint	`gorm:"primary_key,autoIncrement"`
	Title string `json:"caption" bson:"caption,omitempty"`
	Image []string `json:"image", bson:"image,omitempty"`
	Description string `json:"desc" bson:"desc,omitempty"`
}