package mariacustomer

import (
	"neosync/internal/infra/adapter/gorm"
)

type DB struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) *DB {
	return &DB{
		conn: conn,
	}
}
