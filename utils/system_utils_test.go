package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sync"
)


func TestExecCommandBasic(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	status, err := ExecCommand("echo 'tutu'", wg)

	assert.Equal(t, 0, status)
	assert.Equal(t, nil, err)
}

func TestExecCommandWithSleep(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	status, err := ExecCommand("sleep 20; echo tutu", wg)

	assert.Equal(t, 0, status)
	assert.Equal(t, nil, err)
}
