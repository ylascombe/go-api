package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"arc-api/models"
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
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Environment{})
	db.AutoMigrate(&models.EnvironmentAccess{})
	db.AutoMigrate(&models.FeatureTeam{})
	db.AutoMigrate(&models.Membership{})
	db.AutoMigrate(&models.Artefact{})
	db.AutoMigrate(&models.CommonConfig{})
	db.AutoMigrate(&models.AppModule{})
	db.AutoMigrate(&models.Application{})
	db.AutoMigrate(&models.ReactivePlatform{})
	db.AutoMigrate(&models.Manifest{})
}
