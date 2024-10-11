package models

import "time"

type Account struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Number    int64      `json:"number"`
	Balance   int64      `json:"balance"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}
