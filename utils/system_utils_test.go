package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"os/exec"
	"time"
	"io"
	"bytes"
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

	logger := NewLog("/tmp/TestExecCommandListAsynchronously.txt")
	commands := []string{
		"echo start",
		"sleep 4",
		"echo middle",
		"sleep 8",
		"echo 'tutu' >> /tmp/tutu.txt",
		"echo end",
	}
	var reader *bytes.Buffer
	status, err := ExecCommandListAsynchronously(commands)

	assert.Equal(t, 6 , len(status))
	assert.Equal(t, nil, err)


	for i:=0; i<len(status); i++ {
		val := IsTerminated(status[i].Command, logger)
		fmt.Print("log : ")
		reader = Stdout(status[i].Command)

		fmt.Println(reader)
		fmt.Println(i, " : ", val)

	}
	time.Sleep(20 * time.Second)

	assert.Equal(t, "start\nmiddle\nend", reader)
	for i:=0; i<len(status); i++ {
		val := IsTerminated(status[i].Command, logger)

		fmt.Println(i, " : ", val)
	}
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
