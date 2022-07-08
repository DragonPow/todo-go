package db_setup

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"domain"
)

func Init() *gorm.DB {
	url := "postgres://postgres:password@localhost:5432/todo-go"

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Task{}, &domain.Tag{}, &domain.User{})

	return db
}