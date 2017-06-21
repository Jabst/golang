package handlers

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"../db"
	"../Models"
)

func GetTaskByUser(res http.ResponseWriter, req *http.Request){

	id := req.URL.Query().Get("user")

	usr := []models.User{}

	tasks := []models.Task{}

	db := db.ConnectDb()

	db.Where("ID = ?", id).First(&usr)

	db.Model(&usr).Related(&tasks)

	log.Println(tasks)

	jdata, err := json.Marshal(tasks)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jdata)
}

func PostTask(res http.ResponseWriter, req *http.Request){

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var t models.Task

	err = json.Unmarshal(body, &t)

	if err != nil {
		panic(err)
	}

	task := models.Task { Name: t.Name, TaskStatus: t.TaskStatus, UserID: t.UserID }

	db := db.ConnectDb()

	db.NewRecord(task)

	db.Create(&task)

	fmt.Fprint(res, "Criei a seguinte task: %s !\n", t.Name)
}

func DeleteTask(res http.ResponseWriter, req *http.Request){
	
}