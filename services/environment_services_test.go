package services

import (
	"testing"
	"github.com/ylascombe/go-api/models"
	"github.com/stretchr/testify/assert"
)

var (
	envLOCAL models.Environment
	envName = "LOCAL"
)

func TestCreateEnvironment(t *testing.T) {
	result := CreateEnvironment(envName)

	assert.NotNil(t, result)

	assert.Equal(t, envName, result.Name)
	envLOCAL = result
}

func TestGiveAccessTo(t *testing.T) {

	users := ListApiUser()

	result := GiveAccessTo(envLOCAL, users[0])

	assert.NotNil(t, result)
	assert.Equal(t, envName, result.Environment.Name)
}


func TestListEnvironment(t *testing.T) {
	results := ListEnvironment()

	assert.NotNil(t, results)
	assert.Equal(t, 1, len(*results))
}
