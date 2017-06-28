package main

import (
	"log"
	"net/http"
	"os"

	"./Models"
	"./db"
)

func main() {
	var router = NewRouter()

	log.Println(os.Args)

	if argsWithProg := os.Args[1]; argsWithProg == "sync" {

		log.Println("{MESSAGE} : Rebuilding Database...")
		db := db.ConnectDb()

		if db.HasTable(&models.Task{}) {
			db.DropTable(&models.Task{})
			db.CreateTable(&models.Task{})
		} else {
			db.CreateTable(&models.Task{})
		}

		if db.HasTable(&models.User{}) {
			db.DropTable(&models.User{})
			db.CreateTable(&models.User{})
		} else {
			db.CreateTable(&models.User{})
		}

		db.Model(&models.Task{}).AddForeignKey("user_id", "users(ID)", "CASCADE", "CASCADE")

		log.Println("{MESSAGE} : Done")
	}

	log.Fatal(http.ListenAndServe(":8081", router))
}
