package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"os/exec"
)

func TestExecCommandAsynchronously(t *testing.T) {
	logger := NewLog("/tmp/test.txt")

	cmd := exec.Command("echo", "tutu")
	status, err := ExecCommandAsynchronously(cmd, logger)

	fmt.Println(status)
	fmt.Println(err)
	assert.True(t, status.ProcessId > 0)
	assert.Equal(t, nil, err)
}

func TestExecCommandListAsynchronously(t *testing.T) {
	commands := []string{
		"echo start",
		"sleep 4",
		"echo middle",
		"sleep 8",
		"echo end",
	}
	status, err := ExecCommandListAsynchronously(commands)

	assert.Equal(t, 5 , len(status))
	assert.Equal(t, nil, err)

}

func TestExecCommandListAsynchronouslyWhenError(t *testing.T) {
	commands := []string{
		"echo start",
		"curl -Z", // command that produce an error
		"echo end",
	}
	status, err := ExecCommandListAsynchronously(commands)

	assert.Equal(t, 2 , len(status))
	assert.Contains(t, err, "Error when running")
}
