package repository

import (
	"fmt"

	"gorm.io/gorm"

	"project1/util/db"
)

type postgre_repository struct {
	Conn db.Database
}

func newRepository(conn db.Database) *postgre_repository {
	return &postgre_repository{
		Conn: conn,
	}
}

func (t *postgre_repository) GetTransaction(args []interface{}, minNumberLen int) (*gorm.DB, error) {
	var tx *gorm.DB
	if len(args) <= minNumberLen {
		tx = t.Conn.Db
	} else {
		arg0, ok := args[0].(*gorm.DB)
		if ok {
			tx = arg0
		} else {
			return nil, fmt.Errorf("Args tx is needed")
		}
	}
	return tx, nil
}
