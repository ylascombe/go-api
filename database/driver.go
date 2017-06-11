package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDBDriver() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=api dbname=api sslmode=disable password=apipass")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	return db
}

