package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
)

type blog struct {
	title string		`json:"title"`
	content string		`json:"content"`
	category string		`json:"category"`
	tags []string		`json:"tags"`
}

func main() {
	fmt.Println("BLOGGING PLATFORM API")

	// Connecting to the database
	db, err := sql.Open("postgres", "user=postgres password=yourpassword dbname=blogging_platform sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	defer db.Close()
	
	// Starting a HTTP server
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Fatal("Failed to start the HTTP server", err)
	}

	
}