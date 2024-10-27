package api

import (
	"github.com/zhitoo/gobank/storage"
)

type ApiError struct {
	Error string
}

type APIServer struct {
	listenAddr string
	storage    storage.Storage
}

func NewAPIServer(listenAddr string, storage storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}
