package main

import (
	"log"
	"net/http"
)

func main() {
	dsn := "host=localhost user=postgres password=pass dbname=test sslmode=disable"

	repo, _ := NewPostgresRepository(dsn)

	handler := &UserHandler{repo: repo}

	http.HandleFunc("/users", handler.CreateUser)

	log.Println("Start on port: 8080")
	http.ListenAndServe(":8080", nil)
}
