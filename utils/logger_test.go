package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func TestCreateLogger(t *testing.T) {

	filename := "/tmp/test-logger.txt"
	given_line := "Ligne contenu dans le fichier de log"

	actual := NewLog(filename)
	actual.Info(given_line)

	result, err := ioutil.ReadFile(filename)
	assert.Nil(t, err)
	assert.Contains(t, string(result), given_line)
}
