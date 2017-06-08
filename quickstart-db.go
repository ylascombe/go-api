package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"github.com/ylascombe/go-api/models"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=postgres sslmode=disable password=pass")
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	db.AutoMigrate(&models.Artefact{})
	db.AutoMigrate(&models.CommonConfig{})
	db.AutoMigrate(&models.AppModule{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)

	fmt.Println("OK ;-)")
}