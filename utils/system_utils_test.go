package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ylascombe/go-api/models"
)

var responseTasks models.ResponseTask

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
	logger := NewLog("/tmp/test.txt")
	status, err := ExecCommandAsynchronously("echo start", logger)

	assert.True(t, status.ProcessId > 0)
	assert.Equal(t, nil, err)
}

func TestExecCommandListAsynchronously(t *testing.T) {
	commands := [] string {"echo start; sleep 4; echo middle", "sleep 8; echo end"}
	status, err := ExecCommandListAsynchronously(commands)

	assert.Equal(t, 2, len(status))
	assert.Equal(t, nil, err)

}
