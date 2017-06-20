package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name		string		`json:"name"`
	Contact		string		`json:"contact"`

	//Tasks		[]task.Task `gorm:"ForeignKey:UserOwner"`
}

type UserDeletion struct {
	ID		int

	//Tasks		[]task.Task `gorm:"ForeignKey:UserOwner"`
}