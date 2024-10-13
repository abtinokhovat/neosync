package mariaprovider

import (
	"neosync/internal/domain/provider"
	"neosync/internal/infra/adapter/mariadb"
)

type DB struct {
	conn *mariadb.DB
}

func New(conn *mariadb.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func scanProvider(scanner mariadb.Scanner) (provider.Provider, error) {
	p := provider.Provider{}
	err := scanner.Scan(&p.ID, &p.Name, &p.URL)
	return p, err
}
