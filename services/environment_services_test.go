package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"arc-api/database"
	"arc-api/models"
	"testing"
)

type testData struct {
	User *models.User
}

var (
	envName = "SPECIFIC"
)

func TestCreateEnvironment(t *testing.T) {

	tearDown(t)
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
	result, err := GiveAccessTo(*env, *testData.User)

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
	assert.Equal(t, 0, len(*res))

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
	assert.Equal(t, 1, len(*res))

	assert.NotNil(t, envName, (*res)[0].Environment)
	assert.Equal(t, envName, (*res)[0].Environment.Name)
	assert.NotNil(t, (*res)[0].EnvironmentID)
	assert.NotNil(t, (*res)[0].User)
	assert.Equal(t, email, (*res)[0].User.Email)
	assert.Equal(t, firstName, (*res)[0].User.Firstname)
	assert.Equal(t, lastName, (*res)[0].User.Lastname)
	assert.Equal(t, sshPubKey, (*res)[0].User.SshPublicKey)
	assert.Equal(t, pseudo, (*res)[0].User.Pseudo)
	tearDown(t)
}

func TestListAccessForEnvironmentWhenOthers(t *testing.T) {

	tearDown(t)

	// arrange
	otherEnv := "INEXIST"
	testData := setUp(t)
	CreateEnvironment(envName)
	CreateEnvironment(otherEnv)
	AddEnvironmentAccess(testData.User.ID, envName)
	AddEnvironmentAccess(testData.User.ID, otherEnv)

	// act
	res, err := ListAccessForEnvironment(envName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 1, len(*res))

	assert.NotNil(t, (*res)[0])
	assert.NotNil(t, envName, (*res)[0].Environment)
	assert.Equal(t, envName, (*res)[0].Environment.Name)
	assert.NotNil(t, (*res)[0].EnvironmentID)
	assert.NotNil(t, (*res)[0].User)
	assert.Equal(t, email, (*res)[0].User.Email)
	assert.Equal(t, firstName, (*res)[0].User.Firstname)
	assert.Equal(t, lastName, (*res)[0].User.Lastname)
	assert.Equal(t, sshPubKey, (*res)[0].User.SshPublicKey)
	assert.Equal(t, pseudo, (*res)[0].User.Pseudo)
	//tearDown(t)
}

func TestListSshPublicKeyForEnv(t *testing.T) {
	tearDown(t)
	// arrange
	user1, err := CreateUserFromFields("one", "plus", "oplus", "one@corp.com","ssh-id-rsa num1")
	if err != nil {
		panic("db error")
	}

	user2, err := CreateUserFromFields("two", "bar", "tbar", "two@corp.com", "ssh-id-rsa num2")
	if err != nil {
		panic("db error")
	}
	env, _ := CreateEnvironment(envName)
	GiveAccessTo(*env, *user1)
	GiveAccessTo(*env, *user2)

	// act
	res, err := ListSshPublicKeyForEnv(envName)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, "ssh-id-rsa num1 oplus\nssh-id-rsa num2 tbar\n", *res)

	// clean
	tearDown(t)
}

// Force to remove test user
func tearDown(t *testing.T) {

	// remove user in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res2 := db.Exec("delete from environment_accesses").Error
	res1 := db.Exec("delete from users").Error
	res3 := db.Exec("delete from environments").Error

	if res1 != nil || res2 != nil || res3 != nil {
		panic(fmt.Sprintf(
			"db error: users=%s environment_accesses=%s, environments=%s",
			res1, res2, res3))
	}
}

func setUp(t *testing.T) testData {

	user, err := GetUserFromMail(email)

	if err != nil {
		// prerequisites
		user, err = CreateUserFromFields(firstName, lastName, pseudo, email, sshPubKey)
		if err != nil {
			panic("db error")
		}
	}

	// check user has been well created
	assert.Nil(t, err)

	return testData{User: user}
}
