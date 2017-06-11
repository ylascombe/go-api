package services

import (
	"testing"
	"github.com/ylascombe/go-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/ylascombe/go-api/database"
)

var (
	envLOCAL models.Environment
	envName = "SPECIFIC"

	// prerequisites
	user, errPrereq = CreateApiUser("firstName", "lastname", "email@corp.com", "")
)

func TestCreateEnvironment(t *testing.T) {
	result, err := CreateEnvironment(envName)

	assert.NotNil(t, result)
	assert.Nil(t, err)

	assert.Equal(t, envName, result.Name)
	envLOCAL = result

	// remove access in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res := db.Exec("delete from environments where name = ?", envName).Error;
	assert.Nil(t, res)
}

func TestGiveAccessTo(t *testing.T) {

	// check user has been well created
	assert.Nil(t, errPrereq)

	result, err := GiveAccessTo(envLOCAL, user)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, envName, result.Environment.Name)

	// remove access in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res := db.Exec("delete from environment_accesses where api_user_id = ? and environment_id = ?", user.ID, envLOCAL.ID).Error;
	assert.Nil(t, res)
}


func TestListEnvironment(t *testing.T) {
	results, err := ListEnvironment()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(*results))
}


func TestListEnvironmentAccess(t *testing.T) {
	results, err := ListEnvironmentAccesses()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 0, len(results.List))
}

// Force to remove test user
func TestTearDown(t *testing.T) {
	// remove user in order to not change initial state
	db := database.NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res := db.Exec("delete from api_users where email = ?", user.Email).Error;
	assert.Nil(t, res)
}
