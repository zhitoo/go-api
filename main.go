package main

import (
	"log"

	"github.com/zhitoo/gobank/config"
)

func main() {
	storage, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer("localhost:"+config.Envs.Port, storage)
	server.Run()
}
