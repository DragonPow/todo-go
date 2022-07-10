package db

import (
	"log"
	"project1/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Init() *Database {
	url := "postgres://postgres:password@localhost:5432/todo-go"

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Task{}, &domain.Tag{}, &domain.User{})

	return &Database{db: db}
}
