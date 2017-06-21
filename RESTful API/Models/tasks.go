package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model

	Name			string
	TaskStatus		string
	UserID			uint
}