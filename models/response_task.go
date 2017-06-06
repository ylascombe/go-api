package models

import "os/exec"

type ResponseTask struct {
	Command *exec.Cmd
	ProcessId int `yaml:"process_id"`
	TaskCommand string `json:"task_command"`
}
