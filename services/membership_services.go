package services

import (
	"arc-api/models"
	"errors"
	"fmt"
	"arc-api/database"
)

func CreateMembership(user models.User, team models.FeatureTeam) (*models.Membership, error) {

	if ! user.IsValid() {
		return nil, errors.New(fmt.Sprintf("User should be initialized. Got: %s", user))
	}

	if ! team.IsValid() {
		return nil, errors.New(fmt.Sprintf("FeatureTeam should be initialized. Got: %s", team))
	}

	membership, err := models.NewMembership(user, team)

	if err != nil {
		return nil, err
	}

	db := database.NewDBDriver()
	defer db.Close()

	err = db.Create(&membership).Error

	if err != nil {
		return nil, err
	} else {
		res, err := GetMembership(user.ID, team.ID)
		if err == nil {
			return res, nil
		} else {
			return nil, err
		}
	}
}

func CreateMembershipFromIDs(userID uint, featureTeamID uint) (*models.Membership, error) {
	membership := models.Membership{FeatureTeamID: featureTeamID, UserID: userID}

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

		var user models.User
		err = db.First(&user, memberships.List[i].UserID).Error

		if err != nil {
			return nil, err
		}
		memberships.List[i].User = user
		memberships.List[i].FeatureTeam = *featureTeam
	}

	return &memberships, nil
}

func GetMembership(userId uint, featureTeamId uint) (*models.Membership, error) {
	db := database.NewDBDriver()
	defer db.Close()

	membership := models.Membership{UserID: userId, FeatureTeamID: featureTeamId}

	err := db.First(&membership).Error

	if err == nil {
		return &membership, nil
	} else {
		return nil, err
	}
}