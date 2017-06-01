package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sync"
	"fmt"
	"github.com/ylascombe/go-api/models"
)

var responseTasks models.ResponseTask

func TestExecCommandBasic(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	status, err := ExecCommand("echo 'tutu'", wg)

	assert.Equal(t, 0, status)
	assert.Equal(t, nil, err)
}

//func TestExecCommandWithSleep(t *testing.T) {
//	wg := new(sync.WaitGroup)
//	wg.Add(3)
//
//	status, err := ExecCommand("sleep 20; echo tutu", wg)
//
//	assert.Equal(t, 0, status)
//	assert.Equal(t, nil, err)
//}


func TestExecCommandAsynchronously(t *testing.T) {
	commands := [] string {"echo start; sleep 4; echo middle", "sleep 8; echo end"}
	status, err := ExecCommandsAsynchronously(commands)

	assert.True(t, status.ProcessId > 0)
	assert.Equal(t, nil, err)

	responseTasks = status
}



func TestCommandIsAlive(t *testing.T) {
	fmt.Println("is alive ? ", responseTasks)


}



