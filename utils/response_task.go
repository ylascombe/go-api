package utils

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

type ResponseTask struct {
	Command     *exec.Cmd
	ProcessId   int    `yaml:"process_id"`
	TaskCommand string `json:"task_command"`
	Logger      *Logger
}

func IsTerminated(responseTask ResponseTask) bool {
	// TODO add test since it does not work
	return responseTask.Command.ProcessState != nil
}

func (responseTask ResponseTask) StdoutSnapshot() string {

	result, err := ioutil.ReadFile(responseTask.Logger.filepath)

	if err != nil {
		err_msg := fmt.Sprintf("Error when running %s command. Error details: %v\n", responseTask.Command, err)
		fmt.Println(err_msg)
		return "ERROR"
	}
	return string(result)
}
