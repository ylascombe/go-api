package services

import (
	"errors"
	"arc-api/database"
	"arc-api/models"
	"fmt"
	"bytes"
)

func ListEnvironment() (*[]models.Environment, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var environments []models.Environment
	err := db.Find(&environments).Error

	return &environments, err
}

func CreateEnvironment(name string) (*models.Environment, error) {
	env := models.Environment{Name: name}

	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&env).Error

	if err == nil {
		err = db.First(&env).Error
	}

	if err != nil {
		return nil, err
	} else {
		return &env, nil
	}
}

func GiveAccessTo(env models.Environment, user models.User) (*models.EnvironmentAccess, error) {

	envAccess := models.EnvironmentAccess{User: user, UserID:user.ID, Environment: env, EnvironmentID: env.ID}

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

	user, err := GetUser(userID)

	if err != nil {
		return err
	}

	environmentAccess := models.EnvironmentAccess{
		EnvironmentID: environment.ID,
		Environment: *environment,
		UserID: userID,
		User: *user,
		}
	_, err = CreateEnvironmentAccess(&environmentAccess)

	if err != nil {
		return err
	}
	return nil
}

func ListAccessForEnvironment(envName string) (*[]models.EnvironmentAccess, error) {

	var envAccesses []models.EnvironmentAccess

	db := database.NewDBDriver()
	defer db.Close()

	environment, err := GetEnvironmentByName(envName)

	if err != nil {
		return nil, err
	}

	err = db.Model(*environment).Related(&envAccesses).Error

	if err != nil {
		return nil, err
	}

	// load "manually" each item in list
	// TODO manage better error
	for i := 0; i< len(envAccesses); i++ {

		var user models.User

		err = db.First(&user, envAccesses[i].UserID).Error

		if err != nil {
			return nil, err
		}

		envAccesses[i].User = user

		var env models.Environment
		err = db.First(&env, envAccesses[i].EnvironmentID).Error

		if err != nil {
			return nil, err
		}
		envAccesses[i].Environment = env
	}

	return &envAccesses, nil
}

func ListEnvironmentAccesses() (*models.EnvironmentAccesses, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var environmentAccesses []models.EnvironmentAccess
	err := db.Find(&environmentAccesses).Error

	// TODO add test to test empty
	if err != nil {
		return nil, err
	}

	envAccesses := models.EnvironmentAccesses{List: environmentAccesses}
	return &envAccesses, nil
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

func GetEnvironment(id uint) (*models.Environment, error) {
	db := database.NewDBDriver()
	defer db.Close()

	environment := models.Environment{}
	err := db.First(&environment, id).Error

	if err != nil {
		return nil, err
	}

	return &environment, nil
}

func ListSshPublicKeyForEnv(envName string) (*string, error) {

	environment, err := GetEnvironmentByName(envName)

	if err != nil {
		return nil, err
	}

	envAccesses := []models.EnvironmentAccess{}

	db := database.NewDBDriver()
	defer db.Close()


	err = db.Model(environment).Related(&envAccesses).Error

	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer

	for i:=0; i< len(envAccesses); i++ {

		user, err := GetUser(envAccesses[i].UserID)

		if err != nil {
			return nil, err
		}

		buffer.WriteString(user.SshPublicKey + " " + user.Pseudo + "\n")
	}

	result := buffer.String()
	return &result, nil

}