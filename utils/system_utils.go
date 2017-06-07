package utils

import (
	"fmt"
	"os/exec"
	"errors"
	"strings"
)

func ExecCommandListAsynchronously(cmds []string, logger Logger) ([]ResponseTask, error) {
	var commands []ResponseTask
	logger.log.Println("start")
	for i:=0;i<len(cmds);i++ {

		splittedCmd := strings.SplitN(cmds[i]," ",2)

		cmd := exec.Command(splittedCmd[0], splittedCmd[1])
		command, err := ExecCommandAsynchronously(cmd, logger)

		if err != nil {
			return commands, err
		}
		commands = append(commands, command)

		// If we want to be synchronous :
		// commands[command].Wait()

		logger.log.Println(command)
		logger.log.Println("La tache " + cmds[i] + " a terminé, passage à la suivante")

	}
	logger.log.Println("end")

	return commands, nil
}

//func ExecCmdListAsynchronously(cmds []exec.Cmd) ([]models.ResponseTask, error) {
//}

func ExecCommandAsynchronously(command *exec.Cmd, logger Logger) (ResponseTask, error) {

	//logger.log.Println("Prepare to execute command : ", cmd)

	command.Stdout = logger
	err := command.Start()
	//err = command.Wait()
	if err != nil {
		err_msg := fmt.Sprintf("Error when running %s command. Error details: %v\n", command, err)
		fmt.Println("error ----" , command.Path)
		return ResponseTask{}, errors.New(err_msg)
	}

	task := ResponseTask{
		ProcessId: command.Process.Pid,
		TaskCommand: command.Path,
		Command: command,
		Logger: &logger,
	}

	//logger.log.Println("Process : ", command.Process.Pid)
	//logger.log.Println(command.Stdout)
	return task, nil
}

// TODO remove after debug phase
func LaunchTestCommands() {
	logger := NewLog("/tmp/LaunchTestCommands.txt")

	commands := [] string {"echo start; sleep 10; echo middle", "sleep 20; echo end"}
	status, err := ExecCommandListAsynchronously(commands, logger)

	fmt.Println(status)
	fmt.Println(err)
}


func WaitForCompletion(tasks []ResponseTask, logger Logger) []error {
	var errors []error = []error{}

	for i:=0;i<len(tasks);i++ {
		errors = append(errors, tasks[i].Command.Wait())
		logger.log.Println("Task ", tasks[i].TaskCommand, " is terminated")
	}

	return errors
}