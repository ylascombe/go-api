package services

import (
	"errors"
	"fmt"
	"github.com/ylascombe/go-api/database"
	"github.com/ylascombe/go-api/models"
)

func ListEnvironment() (*[]models.Environment, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var environments []models.Environment
	err := db.Find(&environments).Error

	return &environments, err
}

func CreateEnvironment(name string) (models.Environment, error) {
	env := models.Environment{Name: name}

	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&env).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	return env, err
}

func GiveAccessTo(env models.Environment, user models.ApiUser) (*models.EnvironmentAccess, error) {

	envAccess := models.EnvironmentAccess{ApiUser: user, ApiUserID:user.ID, Environment: env, EnvironmentID: env.ID}

	return CreateEnvironmentAccess(&envAccess)
}

func CreateEnvironmentAccess(envAccess *models.EnvironmentAccess) (*models.EnvironmentAccess, error) {

	if !envAccess.IsValid() {
		return nil, errors.New(fmt.Sprintf("Environment should be initialized. Got: %s", envAccess))
	}

	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&envAccess).Error
	return envAccess, err
}

func AddEnvironmentAccess(userID uint, envName string) error {
	environment, err := GetEnvironmentByName(envName)

	if err != nil {
		return err
	}

	apiUser, err := GetApiUser(userID)

	if err != nil {
		return err
	}

	environmentAccess := models.EnvironmentAccess{
		EnvironmentID: environment.ID,
		Environment: *environment,
		ApiUserID: userID,
		ApiUser: *apiUser}
	_, err = CreateEnvironmentAccess(&environmentAccess)

	if err != nil {
		return err
	}
	return nil
}

func ListAccessForEnvironment(envName string) (models.EnvironmentAccesses, error) {

	array := []models.EnvironmentAccess{}
	envAccesses := models.EnvironmentAccesses{List: array}

	db := database.NewDBDriver()
	defer db.Close()

	environment, err := GetEnvironmentByName(envName)

	if err != nil {
		return envAccesses, err
	}

	err = db.Model(environment).Related(&envAccesses.List).Error

	if err != nil {
		return envAccesses, err
	}

	return envAccesses, nil
}

func GetEnvironmentByName(name string) (*models.Environment, error) {
	db := database.NewDBDriver()
	defer db.Close()

	environment := models.Environment{Name: name}
	err := db.First(&environment).Error

	if err != nil {
		return nil, err
	}

	return &environment, nil
}
