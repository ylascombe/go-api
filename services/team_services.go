package services

import (
	"arc-api/database"
	"arc-api/models"
	"errors"
)

func ListFeatureTeams() (*models.FeatureTeams, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var fts []models.FeatureTeam
	err := db.Find(&fts).Error

	if err != nil {
		return nil, err
	}

	featureTeams := models.FeatureTeams{List: fts}
	return &featureTeams, err
}

func CreateFeatureTeam(featureTeam models.FeatureTeam) (*models.FeatureTeam, error) {

	if ! featureTeam.IsValid() {
		return nil, errors.New("Given featureTeam object is not valid")
	}
	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&featureTeam).Error

	if err != nil {
		return nil, err
	} else {
		res, err := GetFeatureTeamFromName(featureTeam.Name)

		if err == nil {
			return res, nil
		} else {
			return nil, err
		}
	}
}


func GetFeatureTeam(ID uint) (*models.FeatureTeam, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var ft models.FeatureTeam
	err := db.First(&ft, "ID = ?", ID).Error

	if err != nil {
		return nil, err
	}

	return &ft, nil
}

func GetFeatureTeamFromName(name string) (*models.FeatureTeam, error) {
	db := database.NewDBDriver()
	defer db.Close()

	var ft models.FeatureTeam
	err := db.First(&ft, "Name = ?", name).Error

	if err != nil {
		return nil, err
	}

	return &ft, nil
}
