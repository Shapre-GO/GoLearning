package main

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Age       int       `json:"id"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at"`
}
