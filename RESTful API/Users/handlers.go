package user

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"../db"
)

func GetAllUsers(res http.ResponseWriter, req *http.Request){

	users := []User{}

	db := db.ConnectDb()

	db.Find(&users)

	jdata, err := json.Marshal(users)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jdata)
}

func GetUserByID(res http.ResponseWriter, req *http.Request){

	id := req.URL.Query().Get("id")

	user := []User{}

	db := db.ConnectDb()

	db.Where("ID = ?", id).First(&user)

	jdata, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jdata)
}

func GetUserByName(res http.ResponseWriter, req *http.Request){
	name := req.URL.Query().Get("name")

	user := []User{}

	db := db.ConnectDb()

	db.Where("Name = ?", name).First(&user)

	jdata, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jdata)
}

func PostUser(res http.ResponseWriter, req *http.Request){

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var u User

	err = json.Unmarshal(body, &u)

	if err != nil {
		panic(err)
	}

	user := User{Name: u.Name, Contact: u.Contact}

	db := db.ConnectDb()

	db.NewRecord(user)

	db.Create(&user)

	fmt.Fprint(res, "Criei o seguinte user: %s !\n", u.Name)
}

func DeleteUser(res http.ResponseWriter, req *http.Request){
	
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var u UserDeletion
	var user User

	err = json.Unmarshal(body, &u)

	if err != nil {
		panic(err)
	}


	db := db.ConnectDb()

	db.Where("ID = ?", u.ID).Find(&user)

	log.Println(user)

	db.Delete(&user)



}