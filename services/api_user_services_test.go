package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ylascombe/go-api/database"
)

func TestCreateApiUser(t *testing.T) {

	email := "john@doe.org"
	user, err := CreateApiUser("john", "doe", email, "ssh-rsa XYZ")

	assert.Nil(t, err)
	assert.NotNil(t, user)

	users, err := ListApiUser()
	assert.Nil(t, err)
	assert.True(t, len(users.List) > 0)

	// remove user in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res := db.Exec("delete from api_users where email = ?", email).Error;
	assert.Nil(t, res)

}

func TestFindApiUsers(t *testing.T) {

	result, err := ListApiUser()
	assert.Nil(t, err)
	assert.True(t, len(result.List) > 0)
}
