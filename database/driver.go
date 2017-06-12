package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ylascombe/go-api/models"
)

type Driver struct {
	DB *gorm.DB
}

func NewDBDriver() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=api dbname=api sslmode=disable password=apipass")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close() DO NOT ADD close here, it has to be done in each calling function
	return db
}

func AutoMigrateDB(db *gorm.DB) {

	// Migrate the schema
	db.AutoMigrate(&models.ApiUser{})
	db.AutoMigrate(&models.Environment{})
	db.AutoMigrate(&models.EnvironmentAccess{})
}
