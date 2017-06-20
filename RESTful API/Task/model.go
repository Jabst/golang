package task

import (
	"github.com/jinzhu/gorm"
	"../Users"
)

type Task struct {
	gorm.Model

	Name			string
	TaskStatus		string
	UserID			uint
	User 			user.User
}