package models

type ResponseTask struct {
	ProcessId int `yaml:"process_id"`
	TaskCommand string `json:"task_command"`
}