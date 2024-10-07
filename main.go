package main

import (
	"log"
)

func main() {
	storage, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":8080", storage)
	server.Run()
}
