package mariaorder

import (
	"neosync/adapter/mariadb"
)

type DB struct {
	conn *mariadb.DB
}

func New(conn *mariadb.DB) *DB {
	return &DB{
		conn: conn,
	}
}
