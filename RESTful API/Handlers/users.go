package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../Models"
	"../db"

	//"../Utils"

	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(res http.ResponseWriter, req *http.Request) {

	users := []models.User{}

	db := db.ConnectDb()

	db.Find(&users)

	jdata, err := json.Marshal(users)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jdata)
}

func GetUserByID(res http.ResponseWriter, req *http.Request) {

	id := req.URL.Query().Get("id")

	user := []models.User{}

	db := db.ConnectDb()

	db.Where("ID = ?", id).First(&user)

	jdata, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jdata)
}

func GetUserByName(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	user := []models.User{}

	db := db.ConnectDb()

	db.Where("Name = ?", name).First(&user)

	jdata, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jdata)
}

//PostUser - Handler for user creation (register)
func PostUser(res http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var u models.UserRegistration

	err = json.Unmarshal(body, &u)

	if err != nil {
		panic(err)
	}

	salt := u.Email + u.Name

	password := []byte(u.Password + salt)	

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	log.Println(string(hashedPassword))

	user := models.User {
					Name 	: u.Name, 
					Contact		: u.Contact,
					Email		: u.Email,
					Password 	: string(hashedPassword[:]),
					Salt 		: salt,
			}



	notErr := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if notErr != nil {
		panic(notErr)
	}

	db := db.ConnectDb()

	db.NewRecord(user)

	db.Create(&user)

	fmt.Fprint(res, "Criei o seguinte user: %s !\n", u.Name)
}

func LoginUser(res http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var ul models.UserLogin

	err = json.Unmarshal(body, &ul)

	user := []models.User{}

	db := db.ConnectDb()

	db.Where("Email = ?", ul.Email).First(&user)

	jdata, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	hashError := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if hashError != nil {
		panic(hashError)

		msg := models.Error{ Message : "Wrong Password!" }

		res.Header().Set("Content-Type", "application/json")
		res.Write(msg)

		return
	}

	res.Header().Set("Content-Type", "application/json")
	cookie := utils.SetToken(ul.Email)
	http.SetCookie(res, &cookie)

	
	res.Write()
	
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var u models.UserDeletion
	var user models.User

	err = json.Unmarshal(body, &u)

	if err != nil {
		panic(err)
	}

	db := db.ConnectDb()

	db.Where("ID = ?", u.ID).Find(&user)

	log.Println(user)

	db.Delete(&user)

}