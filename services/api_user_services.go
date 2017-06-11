package services

import (
	"github.com/ylascombe/go-api/models"
	"github.com/ylascombe/go-api/database"
)

func ListApiUser() []models.ApiUser {

	db := database.NewDBDriver()
	defer db.Close()

	var users []models.ApiUser
	//result := db.Find(&user, "Firstname = ?", "Yohan") // find product with FirstName Yohan
	db.Find(&users) // find product with FirstName Yohan

	return users
}
