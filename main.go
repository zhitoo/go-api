package main

import (
	"log"

	"github.com/zhitoo/gobank/api"
	"github.com/zhitoo/gobank/config"
	"github.com/zhitoo/gobank/storage"
)

func main() {
	storage, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer("localhost:"+config.Envs.Port, storage)
	server.Run()
}
