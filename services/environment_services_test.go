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

	result := GiveAccessTo(envLOCAL, user)

	assert.NotNil(t, result)
	assert.Equal(t, envName, result.Environment.Name)
}


func TestListEnvironment(t *testing.T) {
	results := ListEnvironment()

	assert.NotNil(t, results)
	assert.Equal(t, 1, len(*results))
}
