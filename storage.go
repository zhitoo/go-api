package main

import (
	_ "github.com/lib/pq"
	"github.com/zhitoo/gobank/config"
	"github.com/zhitoo/gobank/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	GetAccountByID(int) (*models.Account, error)
	CreateAccount(FirstName string, LastName string, Number uint64) (*models.Account, error)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStore() (*PostgresStorage, error) {
	dsn := "host=" + config.Envs.DBHost + " user=" + config.Envs.DBUser + " password=" + config.Envs.DBPassword + " dbname=" + config.Envs.DBName + " port=" + config.Envs.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the account schema
	db.AutoMigrate(&models.Account{})

	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (p *PostgresStorage) GetAccountByID(ID int) (*models.Account, error) {
	return &models.Account{}, nil
}

func (p *PostgresStorage) CreateAccount(FirstName string, LastName string, Number uint64) (*models.Account, error) {
	account := models.Account{
		FirstName: FirstName,
		LastName:  LastName,
		Number:    int64(Number),
		Balance:   0,
	}
	result := p.db.Create(&account) // pass pointer of data to Create
	return &account, result.Error
}
