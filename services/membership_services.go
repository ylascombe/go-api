package services

import (
	"github.com/ylascombe/go-api/models"
	"errors"
	"fmt"
	"github.com/ylascombe/go-api/database"
)

func CreateMembership(apiUser models.ApiUser, team models.FeatureTeam) (*models.Membership, error) {

	if ! apiUser.IsValid() {
		return nil, errors.New(fmt.Sprintf("ApiUser should be initialized. Got: %s", apiUser))
	}

	if ! team.IsValid() {
		return nil, errors.New(fmt.Sprintf("FeatureTeam should be initialized. Got: %s", team))
	}

	membership, err := models.NewMembership(apiUser, team)

	if err != nil {
		return nil, err
	}

	db := database.NewDBDriver()
	defer db.Close()

	err = db.Create(&membership).Error

	if err != nil {
		return nil, err
	} else {
		res, err := GetMembership(apiUser.ID, team.ID)
		if err == nil {
			return res, nil
		} else {
			return nil, err
		}
	}
}

func CreateMembershipFromIDs(apiUserID uint, featureTeamID uint) (*models.Membership, error) {
	membership := models.Membership{FeatureTeamID: featureTeamID, ApiUserID: apiUserID}

	db := database.NewDBDriver()
	defer db.Close()

	err := db.Create(&membership).Error

	if err != nil {
		return nil, err
	} else {
		err = db.First(&membership).Error

		if err == nil {
			return &membership, nil
		} else {
			return nil ,err
		}
	}
}

func ListTeamMembers(featureTeamName string) (*models.Memberships, error) {

	array := []models.Membership{}
	memberships := models.Memberships{List: array}

	db := database.NewDBDriver()
	defer db.Close()

	featureTeam, err := GetFeatureTeamFromName(featureTeamName)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Feature team not found. error detail : %s", err))
	}

	err = db.Model(featureTeam).Related(&memberships.List).Error

	if err != nil {
		return nil, err
	}

	// load "manually" each item in list
	// TODO improve error management
	for i := 0; i< len(memberships.List); i++ {

		var apiUser models.ApiUser
		err = db.First(&apiUser, memberships.List[i].ApiUserID).Error

		if err != nil {
			return nil, err
		}
		memberships.List[i].ApiUser = apiUser
		memberships.List[i].FeatureTeam = *featureTeam
	}

	return &memberships, nil
}

func GetMembership(apiUserId uint, featureTeamId uint) (*models.Membership, error) {
	db := database.NewDBDriver()
	defer db.Close()

	membership := models.Membership{ApiUserID: apiUserId, FeatureTeamID: featureTeamId}

	err := db.First(&membership).Error

	if err == nil {
		return &membership, nil
	} else {
		return nil, err
	}
}