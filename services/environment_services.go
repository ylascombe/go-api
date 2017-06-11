package services

import (
	"github.com/ylascombe/go-api/database"
	"github.com/ylascombe/go-api/models"
)

func ListEnvironment() *[]models.Environment {
	db := database.NewDBDriver()
	defer db.Close()

	var environments []models.Environment
	db.Find(&environments)

	return &environments
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

func GiveAccessTo(env models.Environment, user models.ApiUser) models.EnvironmentAccess {

	envAccess := models.EnvironmentAccess{ApiUser: user, Environment: env}

	db := database.NewDBDriver()
	defer db.Close()

	db.Create(&envAccess)
	return envAccess
}

