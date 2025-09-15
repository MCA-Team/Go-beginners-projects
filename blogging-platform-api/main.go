package main

import (
	"fmt"
	"log"
	"net/http"
)

// type blog struct {
// 	title string
// 	content string
// 	category string
// 	tags []string
// }

func main() {
	fmt.Println("BLOGGING PLATFORM API")
	
	// Starting a HTTP server
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Fatal("Failed to start the HTTP server", err)
	}
}