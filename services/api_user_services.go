package services

import (
	"github.com/ylascombe/go-api/database"
	"github.com/ylascombe/go-api/models"
)

func ListApiUser() (*models.ApiUsers, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var users []models.ApiUser
	err := db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	apiUsers := models.ApiUsers{List: users}
	return &apiUsers, err
}

func CreateApiUser(firstName string, lastName string, pseudo string, email string, publicSSHKey string) (*models.ApiUser, error) {


	user, err := models.NewApiUser(firstName, lastName, pseudo, email, publicSSHKey)

	if err != nil {
		return nil, err
	}

	db := database.NewDBDriver()
	defer db.Close()

	err = db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user models.ApiUser) (*models.ApiUser, error) {
	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetApiUser(userID uint) (*models.ApiUser, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var user models.ApiUser
	err := db.First(&user, "ID = ?", userID).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetApiUserFromMail(email string) (*models.ApiUser, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var user models.ApiUser
	err := db.First(&user, "Email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
