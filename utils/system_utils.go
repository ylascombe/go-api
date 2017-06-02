package utils

import (
	"fmt"
	"os/exec"
	"errors"
	"github.com/ylascombe/go-api/models"
)


// TODO is it better to use *sync.WaitGroup as previously ?

func ExecCommandListAsynchronously(cmd []string) ([]models.ResponseTask, error) {

	var commands []models.ResponseTask
	logger := NewLog("/tmp/my.txt")
	fmt.Println("start")
	for i:=0;i<len(cmd);i++ {

		//command, _ := ExecCommandAsynchronously(cmd[i])
		command, _ := ExecCommandAsynchronously(cmd[i], logger)

		commands = append(commands, command)
		//if err != nil {
		//	return nil, err
		//}
		// If we want to be synchronous :
		// commands[command].Wait()

		logger.log.Println(command)
		logger.log.Println("La tache " + cmd[i] + " a terminé, passage à la suivante")

	}
	fmt.Println("end")

	return commands, nil
}

func ExecCommandAsynchronously(cmd string, logger Logger) (models.ResponseTask, error) {

	//logger.log.Println("Prepare to execute command : ", cmd)
	command := exec.Command("sh", "-c", cmd)

	err := command.Start()
	if err != nil {
		err_msg := fmt.Sprintf("Error when running %s command. Error details: %v\n", cmd, err)
		return models.ResponseTask{}, errors.New(err_msg)
	}

	task := models.ResponseTask{
		ProcessId: command.Process.Pid,
		TaskCommand: cmd,
	}

	//logger.log.Println("Process : ", command.Process.Pid)
	//logger.log.Println(command.Stdout)
	return task, nil
}

func LaunchTestCommands() {
	commands := [] string {"echo start; sleep 10; echo middle", "sleep 20; echo end"}
	status, err := ExecCommandListAsynchronously(commands)

	fmt.Println(status)
	fmt.Println(err)
}
