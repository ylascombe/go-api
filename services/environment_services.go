package services

import (
	"github.com/ylascombe/go-api/database"
	"github.com/ylascombe/go-api/models"
	"fmt"
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

func AddEnvironmentAccess(userID uint, envName string) error {
	environment, err := GetEnvironmentFromName(envName)

	// TODO remove println
	fmt.Println("environment.ID", environment.ID)
	fmt.Println("userID", userID)
	if err == nil {
		return err
	}

	fmt.Println(userID)

	environmentAccess := models.EnvironmentAccess{EnvironmentID: environment.ID, ApiUserID: userID}
	_, err = CreateEnvironmentAccess(environmentAccess)

	if err == nil {
		return err
	}
	return nil
}

func ListAccessForEnvironment(name string) (models.EnvironmentAccesses, error) {

	db := database.NewDBDriver()
	defer db.Close()

	environment := models.Environment{Name: name}
	err := db.First(&environment).Error

	if err != nil {
		return models.EnvironmentAccesses{}, err
	}

	var environmentAccess []models.EnvironmentAccess
	err = db.Model(environment).Related(&environmentAccess).Error

	if err != nil {
		return models.EnvironmentAccesses{}, err
	}

	array := []models.EnvironmentAccess{}
	return models.EnvironmentAccesses{List: array}, nil
}

func GetEnvironmentFromName(name string) (models.Environment, error) {
	db := database.NewDBDriver()
	defer db.Close()

	environment := models.Environment{Name: name}
	err := db.First(&environment).Error;

	return environment, err
}

