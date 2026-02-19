package main

import (
	"log"
	"net/http"

	"web-go/db"
	"web-go/handler"
)

func main() {
	db.InitNeo4j()
	defer db.Driver.Close(nil)

	http.HandleFunc("/users", handler.GetUsers)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
