package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func TestCreateLogger(t *testing.T) {

	filename := "/tmp/test-logger.txt"
	given_line := "Ligne contenue dans le fichier de log"

	actual := NewLog(filename)
	actual.Info(given_line)

	result, err := ioutil.ReadFile(filename)
	assert.Nil(t, err)
	assert.Contains(t, string(result), given_line)

	assert.Equal(t, filename, actual.GetLogpath())
}

func TestWrite(t *testing.T) {
	filename := "/tmp/TestWrite.txt"
	given_line := "start"
	byteArray := []byte{115, 116, 97, 114, 116, 10}
	actual := NewLog(filename)
	actual.Write(byteArray)

	result, err := ioutil.ReadFile(filename)
	assert.Nil(t, err)
	assert.Contains(t, string(result), given_line)
}