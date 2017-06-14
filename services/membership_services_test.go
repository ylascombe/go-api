package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/ylascombe/go-api/database"
	"github.com/ylascombe/go-api/models"
)

func TestCreateMembership(t *testing.T) {

	tearDownMembership(t)
	// arrange
	featureTeam, _ := models.NewFeatureTeam(ftName)
	CreateFeatureTeam(*featureTeam)
	featureTeam, _ = GetFeatureTeamFromName(ftName)

	CreateApiUser(firstName, lastName, pseudo, email, sshPubKey)
	apiUser, _ := GetApiUserFromMail(email)

	// act
	membership, err := CreateMembership(*apiUser, *featureTeam)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, membership)
	//assert.Equal(t, 1, len(res.List))

	//clean
	tearDownMembership(t)
}

func TestListTeamsMembersInexistantTeam(t *testing.T) {
	// arrange

	// act
	res, err := ListTeamMembers("inexistant")

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestListTeamsMembersEmptyList(t *testing.T) {
	// arrange
	featureTeam, _ := models.NewFeatureTeam(ftName)
	CreateFeatureTeam(*featureTeam)

	// act
	res, err := ListTeamMembers(ftName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 0, len(res.List))

	//clean
	tearDownMembership(t)
}

func TestListTeamsMembersWhenNotEmpty(t *testing.T) {
	// in case of existing values
	tearDownMembership(t)

	// arrange
	featureTeam, _ := models.NewFeatureTeam(ftName)
	CreateFeatureTeam(*featureTeam)

	// request it to avoid to create again feature team
	featureTeam, _ = GetFeatureTeamFromName(ftName)

	apiUser, _ := CreateApiUser(firstName, lastName, pseudo, email, sshPubKey)

	CreateMembership(*apiUser, *featureTeam)

	// act
	res, err := ListTeamMembers(ftName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 1, len(res.List))
	assert.Equal(t, ftName, res.List[0].FeatureTeam.Name)
	assert.Equal(t, pseudo, res.List[0].ApiUser.Pseudo)

	//clean
	tearDownMembership(t)
}

func TestCreateMembershipFromIDs(t *testing.T) {
	// arrange
	featureTeam, _ := models.NewFeatureTeam(ftName)
	CreateFeatureTeam(*featureTeam)
	CreateApiUser(firstName, lastName, pseudo, email, sshPubKey)

	featureTeam, _ = GetFeatureTeamFromName(ftName)
	apiUser, _ := GetApiUserFromMail(email)

	// act
	membership, err := CreateMembershipFromIDs(apiUser.ID, featureTeam.ID)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, membership)

	//clean
	tearDownMembership(t)
}

func TestCreateMembershipAlreadyExists(t *testing.T) {
	// arrange
	featureTeam, _ := models.NewFeatureTeam(ftName)
	CreateFeatureTeam(*featureTeam)
	CreateApiUser(firstName, lastName, pseudo, email, sshPubKey)

	featureTeam, _ = GetFeatureTeamFromName(ftName)
	apiUser, _ := GetApiUserFromMail(email)

	// create a first time
	membership, err := CreateMembershipFromIDs(apiUser.ID, featureTeam.ID)

	// act
	membership, err = CreateMembershipFromIDs(apiUser.ID, featureTeam.ID)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, membership)



	//clean
	tearDownMembership(t)
}

func tearDownMembership(t *testing.T) {

	db := database.NewDBDriver()
	defer db.Close()

	res2 := db.Exec("delete from memberships").Error
	res1 := db.Exec("delete from feature_teams").Error
	res3 := db.Exec("delete from api_users").Error

	if res1 != nil || res2 != nil {
		panic(fmt.Sprintf(
			"db error: feature_teams=%s memberships=%s api_users=%s",
			res1, res2, res3))
	}
}