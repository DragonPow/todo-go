package db

import (
	"log"
	"project1/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func Init() *Database {
	url := "postgres://postgres:111200@localhost:5432/todo-go"

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Task{}, &domain.Tag{}, &domain.User{})

	return &Database{Db: db}
}

// func (db *Database) Transaction(fc func(params *Database) error, opts ...*sql.TxOptions) (err error) {

// 	return db.Db.Transaction(func(tx *gorm.DB) error { return fc(db) }, opts...)
// }
