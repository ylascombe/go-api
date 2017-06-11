package services

import (
	"github.com/ylascombe/go-api/models"
	"github.com/ylascombe/go-api/database"
)

func ListApiUser() ([]models.ApiUser, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var users []models.ApiUser
	//result := db.Find(&user, "Firstname = ?", "Yohan") // find product with FirstName Yohan
	err := db.Find(&users).Error;

	return users, err
}

func CreateApiUser(firstName string, lastName string, email string, publicSSHKey string) (models.ApiUser, error) {
	user := models.ApiUser{Lastname: lastName, Firstname: firstName, Email: email, SshPublicKey: publicSSHKey}

	db := database.NewDBDriver()
	defer db.Close()

	err:= db.Create(&user).Error;
	//if err != nil {
	//	fmt.Println(err)
	//}

	return user, err
}
