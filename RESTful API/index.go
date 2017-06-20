package main

import (
	"log"
	"net/http"
)

func main() {
	var router = NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}