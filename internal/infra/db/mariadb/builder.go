package mariadb

import (
	"neosync/internal/infra/adapter/mariadb"
	"neosync/internal/infra/db/mariadb/mariaorder"
	"neosync/internal/infra/db/mariadb/mariaprovider"
)

type Databases struct {
	Order    *mariaorder.DB
	Provider *mariaprovider.DB
}

func Builder(adapter *mariadb.DB) *Databases {
	return &Databases{
		Order:    mariaorder.New(adapter),
		Provider: mariaprovider.New(adapter),
	}
}
