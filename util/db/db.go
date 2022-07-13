package db

import (
	"log"
	"os"
	"project1/domain"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db *gorm.DB
}

func Init() *Database {
	url := "postgres://postgres:111200@localhost:5432/todo-go"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Task{}, &domain.Tag{}, &domain.User{})

	return &Database{Db: db}
}

// func (db *Database) Transaction(fc func(params *Database) error, opts ...*sql.TxOptions) (err error) {

// 	return db.Db.Transaction(func(tx *gorm.DB) error { return fc(db) }, opts...)
// }
