package utils

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestExecCommandAsynchronously(t *testing.T) {
	logger := NewLog("/tmp/test.txt")

	cmd := exec.Command("echo", "tutu")
	status, err := ExecCommandAsynchronously(cmd, logger)

	assert.True(t, status.ProcessId > 0)
	assert.Equal(t, nil, err)

	assert.Contains(t, status.StdoutSnapshot(), "tutu")
}

func TestExecCommandListAsynchronously(t *testing.T) {

	logger := NewLog("/tmp/TestExecCommandListAsynchronously.txt")
	commands := []string{
		"echo start",
		"sleep 4",
		"echo 'ligne_de_test'",
		"sleep 8",
		"echo 'tutu' >> /tmp/tutu.txt",
		"echo end",
	}
	status, err := ExecCommandListAsynchronously(commands, logger)

	assert.Equal(t, 6, len(status))
	assert.Equal(t, nil, err)

	for i := 0; i < len(status); i++ {
		assert.False(t, IsTerminated(status[i]))
	}

	WaitForCompletion(status, logger)
	for i := 0; i < len(status); i++ {
		assert.True(t, IsTerminated(status[i]))
	}
	stdout := status[0].StdoutSnapshot
	assert.Contains(t, stdout, "ligne_de_test")
}

func TestExecCommandListAsynchronouslyWhenError(t *testing.T) {

	logger := NewLog("/tmp/TestExecCommandListAsynchronouslyWhenError.txt")

	commands := []string{
		"echo start",
		"curl -Z", // command that produce an error
		"echo end",
	}
	status, _ := ExecCommandListAsynchronously(commands, logger)

	assert.Equal(t, 3, len(status))

	errors := WaitForCompletion(status, logger)
	assert.NotNil(t, errors)
	assert.Equal(t, 3, len(errors))

	assert.Nil(t, errors[0])
	assert.NotNil(t, errors[1])
	//assert.Equal(t,  reflect.Type(exec.ExitError{}), string(reflect.TypeOf(errors[1]).Kind()))
	assert.Nil(t, errors[2])
}
