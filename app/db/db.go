package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase(dia, uri string) *gorm.DB {
	var dialector gorm.Dialector
	switch dia {
	case "sqlite":
		dialector = sqlite.Open(uri)
	case "mysql":
		dialector = mysql.Open(uri)
	case "postgres":
		dialector = postgres.Open(uri)
	}

	database, _ := gorm.Open(dialector)
	return database
}
