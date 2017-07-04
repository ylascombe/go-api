package services

import (
	"arc-api/database"
	"arc-api/models"
	"fmt"
)

func ListUser() (*models.Users, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var users []models.User
	err := db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	res := models.Users{List: users}
	return &res, err
}

func CreateUserFromFields(firstName string, lastName string, pseudo string, email string, publicSSHKey string) (*models.User, error) {

	user, err := models.NewUser(firstName, lastName, pseudo, email, publicSSHKey)

	if err != nil || ! user.IsValid() {
		fmt.Println("user is not valid")
		return nil, err
	}

	db := database.NewDBDriver()
	defer db.Close()

	err = db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	res, err := GetUserFromMail(email)

	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func CreateUser(user models.User) (*models.User, error) {
	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userID uint) (*models.User, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var user models.User
	err := db.First(&user, "ID = ?", userID).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserFromMail(email string) (*models.User, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var user models.User
	err := db.First(&user, "Email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
