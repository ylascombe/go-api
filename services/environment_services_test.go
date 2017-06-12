package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/ylascombe/go-api/database"
	"github.com/ylascombe/go-api/models"
	"testing"
)

type testData struct {
	User *models.ApiUser
}

var (
	envName = "SPECIFIC"
)

func TestCreateEnvironment(t *testing.T) {

	// Arrange
	setUp(t)
	db := database.NewDBDriver()
	defer db.Close()

	var countAfter int
	var countBefore int
	db.Model(&models.Environment{}).Count(&countBefore)

	// Act
	result, err := CreateEnvironment(envName)

	// Assert
	db.Model(&models.Environment{}).Count(&countAfter)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	assert.Equal(t, envName, result.Name)
	assert.True(t, countAfter == countBefore+1)

	tearDown(t)
}

func TestGiveAccessToInvalidEnvironment(t *testing.T) {

	// arrange
	testData := setUp(t)
	db := database.NewDBDriver()
	defer db.Close()

	var countAfter int
	var countBefore int

	db.Model(&models.EnvironmentAccess{}).Count(&countBefore)

	// act
	result, err := GiveAccessTo(models.Environment{}, *testData.User)

	// assert
	assert.Nil(t, result)
	assert.NotNil(t, err)

	db.Model(&models.EnvironmentAccess{}).Count(&countAfter)
	assert.True(t, countAfter == countBefore)

	tearDown(t)
}

func TestGiveAccessTo(t *testing.T) {

	// arrange
	testData := setUp(t)
	db := database.NewDBDriver()
	defer db.Close()

	var countAfter int
	var countBefore int

	env, _ := CreateEnvironment(envName)
	db.Model(&models.EnvironmentAccess{}).Count(&countBefore)

	// act
	result, err := GiveAccessTo(env, *testData.User)

	// assert
	assert.NotNil(t, result)
	assert.Nil(t, err)

	db.Model(&models.EnvironmentAccess{}).Count(&countAfter)
	assert.Equal(t, countBefore+1, countAfter)

	tearDown(t)
}


func TestListEnvironmentWhenEmpty(t *testing.T) {
	// arrange

	// act
	results, err := ListEnvironment()

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 0, len(*results))
}

func TestListEnvironmentWhenOne(t *testing.T) {
	// arrange
	setUp(t)
	CreateEnvironment("TEST")

	// act
	results, err := ListEnvironment()

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(*results))

	tearDown(t)
}

func TestAddEnvironmentAccessWhenEnvironmentIsAbsent(t *testing.T) {

	// arrange
	var countBefore int
	var countAfter int

	testData := setUp(t)
	db := database.NewDBDriver()
	defer db.Close()
	db.Model(models.EnvironmentAccess{}).Count(&countBefore)

	// act
	err := AddEnvironmentAccess(testData.User.ID, envName)

	// assert
	assert.NotNil(t, err)
	db.Model(models.EnvironmentAccess{}).Count(&countAfter)
	assert.Equal(t, countBefore, countAfter)

	tearDown(t)
}

func TestAddEnvironmentAccessWhenUserIsAbsent(t *testing.T) {

	// arrange
	var countBefore int
	var countAfter int

	testData := setUp(t)
	db := database.NewDBDriver()
	defer db.Close()

	CreateEnvironment(envName)

	// act
	err := AddEnvironmentAccess(testData.User.ID + 42, envName)

	// assert
	assert.NotNil(t, err)
	db.Model(models.EnvironmentAccess{}).Count(&countAfter)
	assert.Equal(t, countBefore, countAfter)

	tearDown(t)
}

func TestAddEnvironmentAccess(t *testing.T) {

	// arrange
	var countBefore int
	var countAfter int

	testData := setUp(t)
	db := database.NewDBDriver()
	defer db.Close()
	db.Model(models.EnvironmentAccess{}).Count(&countBefore)

	CreateEnvironment(envName)

	// act
	err := AddEnvironmentAccess(testData.User.ID, envName)

	// assert
	assert.Nil(t, err)
	db.Model(models.EnvironmentAccess{}).Count(&countAfter)
	assert.Equal(t, countBefore+1, countAfter)

	tearDown(t)
}

func TestListEnvironmentAccess(t *testing.T) {
	results, err := ListEnvironmentAccesses()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 0, len(results.List))
}

func TestGetEnvironment(t *testing.T) {

	// arrange
	setUp(t)
	tmpEnv, _ := CreateEnvironment(envName)

	// act
	res, err := GetEnvironment(tmpEnv.ID)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, envName, res.Name)

	tearDown(t)
}

func TestGetEnvironmentByName(t *testing.T) {
	// arrange
	setUp(t)
	CreateEnvironment(envName)

	// act
	res, err := GetEnvironmentByName(envName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, envName, res.Name)

	tearDown(t)

}

func TestGetEnvironmentByNameWhenEnvironmentIsAbsent(t *testing.T) {
	// arrange
	setUp(t)

	// act
	res, err := GetEnvironmentByName(envName)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, res)

	tearDown(t)
}

func TestListAccessForEnvironmentWhenEmpty(t *testing.T) {
	// arrange
	setUp(t)
	CreateEnvironment(envName)

	// act
	res, err := ListAccessForEnvironment(envName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 0, len(res.List))

	tearDown(t)
}

func TestListAccessForEnvironmentWhenOne(t *testing.T) {
	// arrange
	testData := setUp(t)
	CreateEnvironment(envName)
	err := AddEnvironmentAccess(testData.User.ID, envName)

	// act
	res, err := ListAccessForEnvironment(envName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 1, len(res.List))

	assert.NotNil(t, envName, res.List[0].Environment)
	assert.Equal(t, envName, res.List[0].Environment.Name)
	assert.NotNil(t, res.List[0].EnvironmentID)
	assert.NotNil(t, res.List[0].ApiUser)
	assert.Equal(t, email, res.List[0].ApiUser.Email)
	assert.Equal(t, firstName, res.List[0].ApiUser.Firstname)
	assert.Equal(t, lastName, res.List[0].ApiUser.Lastname)
	assert.Equal(t, sshPubKey, res.List[0].ApiUser.SshPublicKey)
	tearDown(t)
}

// Force to remove test user
func tearDown(t *testing.T) {

	// remove user in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res1 := db.Exec("delete from api_users").Error
	res2 := db.Exec("delete from environment_accesses").Error
	res3 := db.Exec("delete from environments").Error

	if res1 != nil || res2 != nil || res3 != nil {
		panic(fmt.Sprintf(
			"db error: api_users=%s environment_accesses=%s, environments=%s",
			res1, res2, res3))
	}
}

func setUp(t *testing.T) testData {

	user, err := GetApiUserFromMail(email)

	if err != nil {
		// prerequisites
		user, err = CreateApiUser(firstName, lastName, email, sshPubKey)
		if err != nil {
			panic("db error")
		}
	}

	// check user has been well created
	assert.Nil(t, err)

	return testData{User: user}
}
