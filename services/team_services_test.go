package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"arc-api/database"
	"arc-api/models"
	"fmt"
)

var (
	ftName = "ftref"
)

func TestCreateFeatureTeam(t *testing.T) {
	// arrange
	var countBefore int
	var countAfter int

	featureTeam := models.FeatureTeam{Name: ftName}

	db := database.NewDBDriver()
	defer db.Close()
	db.Model(&models.FeatureTeam{}).Count(&countBefore)

	// act
	resultFT, err := CreateFeatureTeam(featureTeam)

	// assert
	db.Model(&models.FeatureTeam{}).Count(&countAfter)

	assert.Nil(t, err)
	assert.NotNil(t, resultFT)
	assert.Equal(t, countBefore+1, countAfter)
	assert.Equal(t, ftName, resultFT.Name)

	tearDownFeatureTeam(t)
}


func TestListFeatureTeamsWhenNone(t *testing.T) {
	// arrange

	// act
	featuresTeams, err := ListFeatureTeams()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 0, len(featuresTeams.List))
}

func TestListFeatureTeamsWhenOne(t *testing.T) {
	// arrange
	featureTeam := models.FeatureTeam{Name: ftName}
	CreateFeatureTeam(featureTeam)

	// act
	featuresTeams, err := ListFeatureTeams()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(featuresTeams.List))

	assert.Equal(t, ftName, featuresTeams.List[0].Name)

	// clean
	tearDownFeatureTeam(t)
}

func TestGetFeatureTeamWhenAbsent(t *testing.T) {

	// act
	res, err := GetFeatureTeam(9999)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestGetFeatureTeamWhenPresent(t *testing.T) {

	// arrange
	featureTeam := models.FeatureTeam{Name: ftName}
	CreateFeatureTeam(featureTeam)
	list, _ := ListFeatureTeams()

	// act
	res, err := GetFeatureTeam(list.List[0].ID)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, list.List[0].Name, res.Name)
	assert.Equal(t, list.List[0].ID, res.ID)
	assert.Equal(t, list.List[0].CreatedAt, res.CreatedAt)

	// clean
	tearDownFeatureTeam(t)
}

func TestGetFeatureTeamFromNameWhenExist(t *testing.T) {

	// arrange
	featureTeam := models.FeatureTeam{Name: ftName}
	CreateFeatureTeam(featureTeam)

	// act
	res, err := GetFeatureTeamFromName(ftName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, ftName, res.Name)

	// clean
	tearDownFeatureTeam(t)
}

func TestGetFeatureTeamFromNameWhenAbsent(t *testing.T) {

	// arrange

	// act
	res, err := GetFeatureTeamFromName(ftName)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, res)
}


func tearDownFeatureTeam(t *testing.T) {

	// remove feature teams in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res1 := db.Exec("delete from feature_teams").Error

	if res1 != nil {
		panic(fmt.Sprintf(
			"db error: feature_teams=%s",
			res1))
	}
}