package main

import (
	"log"
	"net/http"
	"os"
	"./db"
	"./Models"
)

func main() {
	var router = NewRouter()

	log.Println(os.Args)

	if argsWithProg := os.Args[1] ; argsWithProg == "sync"{

		log.Println("{MESSAGE} : Rebuilding Database...")
		db := db.ConnectDb()

		if db.HasTable(&models.User{}){
			db.DropTable(&models.User{})
			db.CreateTable(&models.User{})
		} else {
			db.CreateTable(&models.User{})
		}

		if db.HasTable(&models.Task{}){
			db.DropTable(&models.Task{})
			db.CreateTable(&models.Task{})	
		} else {
			db.CreateTable(&models.Task{})	
		}

		db.Model(&models.Task{}).AddForeignKey("user_id", "users(ID)", "CASCADE", "CASCADE")



		log.Println("{MESSAGE} : Done")
	}

	log.Fatal(http.ListenAndServe(":8081", router))
}