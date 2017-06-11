package services

import (
	"testing"
	"github.com/ylascombe/go-api/models"
	"github.com/stretchr/testify/assert"
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
}

func TestGiveAccessTo(t *testing.T) {

	// check user has been well created
	assert.Nil(t, errPrereq)

	result, err := GiveAccessTo(envLOCAL, user)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, envName, result.Environment.Name)
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
