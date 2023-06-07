package main

import (
	"log"

	a "github.com/jobquestvault/platform-go-challenge/internal/app"
)

func main() {
	server := a.NewServer(8080)

	log.Println("Server starting...")
	err := server.Start()
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
