package services

import (
	"github.com/stretchr/testify/assert"
	"arc-api/database"
	"testing"
	"fmt"
	"arc-api/models"
)

var (
	email = "john@doe.org"
	firstName = "john"
	lastName = "doe"
	pseudo = "jdoe"
	sshPubKey = "ssh-rsa XYZ"
)

func TestCreateUser(t *testing.T) {

	// arrange
	var countBefore int
	var countAfter int

	db := database.NewDBDriver()
	defer db.Close()
	db.Model(&models.User{}).Count(&countBefore)

	// act
	user, err := CreateUserFromFields(firstName, lastName, pseudo, email, sshPubKey)

	// assert
	db.Model(&models.User{}).Count(&countAfter)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, countBefore+1, countAfter)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, firstName, user.Firstname)
	assert.Equal(t, lastName, user.Lastname)
	assert.Equal(t, sshPubKey, user.SshPublicKey)

	tearDownUser(t)
}

func TestListUserWhenNoUser(t *testing.T) {

	// act
	users, err := ListUser()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 0, len(users.List))
}

func TestListUserWhenOneUser(t *testing.T) {

	// arrange
	CreateUserFromFields(firstName, lastName, pseudo, email, sshPubKey)

	// act
	users, err := ListUser()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users.List))

	assert.Equal(t, email, users.List[0].Email)

	// clean
	tearDownUser(t)
}

func TestGetUserWhenUserIsAbsent(t *testing.T) {

	// act
	user, err := GetUser(99999999)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func tearDownUser(t *testing.T) {

	// remove user in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res1 := db.Exec("delete from users").Error

	if res1 != nil {
		panic(fmt.Sprintf(
			"db error: users=%s",
			res1))
	}
}

