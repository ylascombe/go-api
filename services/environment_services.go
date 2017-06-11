package services

import (
	"github.com/ylascombe/go-api/database"
	"github.com/ylascombe/go-api/models"
)

func ListEnvironment() (*[]models.Environment, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var environments []models.Environment
	err := db.Find(&environments).Error;

	return &environments, err
}

func CreateEnvironment(name string) (models.Environment, error) {
	env := models.Environment{Name: name}

	db := database.NewDBDriver()
	defer db.Close()

	err:= db.Create(&env).Error;
	//if err != nil {
	//	fmt.Println(err)
	//}

	return env, err
}

func GiveAccessTo(env models.Environment, user models.ApiUser) (models.EnvironmentAccess, error) {

	envAccess := models.EnvironmentAccess{ApiUser: user, Environment: env}

	return CreateEnvironmentAccess(envAccess)
}

func CreateEnvironmentAccess(envAccess models.EnvironmentAccess) (models.EnvironmentAccess, error) {

	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&envAccess).Error;
	return envAccess, err
}
