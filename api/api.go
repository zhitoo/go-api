package api

import (
	"github.com/zhitoo/gobank/requests"
	"github.com/zhitoo/gobank/storage"
)

type ApiError struct {
	Error string
}

type APIServer struct {
	listenAddr string
	storage    storage.Storage
	validator  *requests.Validator
}

func NewAPIServer(listenAddr string, storage storage.Storage, validator *requests.Validator) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
		validator:  validator,
	}
}
