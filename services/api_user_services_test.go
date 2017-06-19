package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/ylascombe/go-api/database"
	"testing"
	"fmt"
	"github.com/ylascombe/go-api/models"
)

var (
	email = "john@doe.org"
	firstName = "john"
	lastName = "doe"
	pseudo = "jdoe"
	sshPubKey = "ssh-rsa XYZ"
)

func TestCreateApiUser(t *testing.T) {

	// arrange
	var countBefore int
	var countAfter int

	db := database.NewDBDriver()
	defer db.Close()
	db.Model(&models.ApiUser{}).Count(&countBefore)

	// act
	user, err := CreateApiUser(firstName, lastName, pseudo, email, sshPubKey)

	// assert
	db.Model(&models.ApiUser{}).Count(&countAfter)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, countBefore+1, countAfter)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, firstName, user.Firstname)
	assert.Equal(t, lastName, user.Lastname)
	assert.Equal(t, sshPubKey, user.SshPublicKey)

	tearDownApiUser(t)
}

func TestListApiUserWhenNoUser(t *testing.T) {

	// act
	users, err := ListApiUser()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 0, len(users.List))
}

func TestListApiUserWhenOneUser(t *testing.T) {

	// arrange
	CreateApiUser(firstName, lastName, pseudo, email, sshPubKey)

	// act
	users, err := ListApiUser()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users.List))

	assert.Equal(t, email, users.List[0].Email)

	// clean
	tearDownApiUser(t)
}

func TestGetApiUserWhenUserIsAbsent(t *testing.T) {

	// act
	user, err := GetApiUser(99999999)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func tearDownApiUser(t *testing.T) {

	// remove user in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res1 := db.Exec("delete from api_users").Error

	if res1 != nil {
		panic(fmt.Sprintf(
			"db error: api_users=%s",
			res1))
	}
}

