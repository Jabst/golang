package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name		string		`json:"name"`
	Contact		string		`json:"contact"`

	Tasks		[]Task      `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
}

type UserDeletion struct {
	ID		int

	//Tasks		[]task.Task `gorm:"ForeignKey:UserOwner"`
}

