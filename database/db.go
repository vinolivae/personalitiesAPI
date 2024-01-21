//gorm setup

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// iniciando variáveis.
var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	connParams := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connParams))
	if err != nil {
		panic(err.Error())
	}
}
