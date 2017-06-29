package models

import (
	"github.com/jinzhu/gorm"
)

//User  - Model for User in the database
type User struct {
	gorm.Model

	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`

	Tasks []Task `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
}

//UserRegistration - Model for User when registered
type UserRegistration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name	 string `json:"name"`
	Contact	 string `json:"contact"`
}

//UserDeletion  - Model for Deleting an user used to format the JSON respsone from the Web Client
type UserDeletion struct {
	ID int

	//Tasks		[]task.Task `gorm:"ForeignKey:UserOwner"`
}

//UserLogin - Model for carrying the plain-text password and the email
type UserLogin struct {
	Email    	string
	Password 	string
}

type UserLoggedIn struct {
	Token 		string 	`json:"token"`
}
