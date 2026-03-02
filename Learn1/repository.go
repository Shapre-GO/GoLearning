package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(u *User) error
	GetAll() ([]User, error)
}

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(dsn string) (*PostgresRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Error DB %w", err)
	}
	db.AutoMigrate(&User{})
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Create(u *User) error {
	return r.db.Create(u).Error
}
func (r *PostgresRepository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}
