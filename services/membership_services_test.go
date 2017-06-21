package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"arc-api/database"
	"arc-api/models"
)

func TestCreateMembership(t *testing.T) {

	tearDownMembership(t)
	// arrange
	featureTeam, _ := models.NewFeatureTeam(ftName, gitlabUrl, groupId)
	CreateFeatureTeam(*featureTeam)
	featureTeam, _ = GetFeatureTeamFromName(ftName)

	CreateUserFromFields(firstName, lastName, pseudo, email, sshPubKey)
	user, _ := GetUserFromMail(email)

	// act
	membership, err := CreateMembership(*user, *featureTeam)

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
	featureTeam, _ := models.NewFeatureTeam(ftName, gitlabUrl, groupId)
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
	featureTeam, _ := models.NewFeatureTeam(ftName, gitlabUrl, groupId)
	CreateFeatureTeam(*featureTeam)

	// request it to avoid to create again feature team
	featureTeam, _ = GetFeatureTeamFromName(ftName)

	user, _ := CreateUserFromFields(firstName, lastName, pseudo, email, sshPubKey)

	CreateMembership(*user, *featureTeam)

	// act
	res, err := ListTeamMembers(ftName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 1, len(res.List))
	assert.Equal(t, ftName, res.List[0].FeatureTeam.Name)
	assert.Equal(t, pseudo, res.List[0].User.Pseudo)

	//clean
	tearDownMembership(t)
}

func TestCreateMembershipFromIDs(t *testing.T) {
	// arrange
	featureTeam, _ := models.NewFeatureTeam(ftName, gitlabUrl, groupId)
	CreateFeatureTeam(*featureTeam)
	CreateUserFromFields(firstName, lastName, pseudo, email, sshPubKey)

	featureTeam, _ = GetFeatureTeamFromName(ftName)
	user, _ := GetUserFromMail(email)

	// act
	membership, err := CreateMembershipFromIDs(user.ID, featureTeam.ID)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, membership)

	//clean
	tearDownMembership(t)
}

func TestCreateMembershipAlreadyExists(t *testing.T) {
	// arrange
	featureTeam, _ := models.NewFeatureTeam(ftName, gitlabUrl, groupId)
	CreateFeatureTeam(*featureTeam)
	CreateUserFromFields(firstName, lastName, pseudo, email, sshPubKey)

	featureTeam, _ = GetFeatureTeamFromName(ftName)
	user, _ := GetUserFromMail(email)

	// create a first time
	membership, err := CreateMembershipFromIDs(user.ID, featureTeam.ID)

	// act
	membership, err = CreateMembershipFromIDs(user.ID, featureTeam.ID)

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
	res3 := db.Exec("delete from users").Error

	if res1 != nil || res2 != nil {
		panic(fmt.Sprintf(
			"db error: feature_teams=%s memberships=%s users=%s",
			res1, res2, res3))
	}
}