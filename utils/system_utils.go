package utils

import (
	"fmt"
	"os/exec"
	"errors"
	"github.com/ylascombe/go-api/models"
	"strings"
	"io/ioutil"
)


// TODO is it better to use *sync.WaitGroup as previously ?

func ExecCommandListAsynchronously(cmds []string) ([]models.ResponseTask, error) {
	var commands []models.ResponseTask
	logger := NewLog("/tmp/my.txt")
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

func ExecCommandAsynchronously(command *exec.Cmd, logger Logger) (models.ResponseTask, error) {

	//logger.log.Println("Prepare to execute command : ", cmd)

	err := command.Start()
	//err = command.Wait()
	if err != nil {
		err_msg := fmt.Sprintf("Error when running %s command. Error details: %v\n", command, err)
		fmt.Println("error ----" , command.Path)
		return models.ResponseTask{}, errors.New(err_msg)
	}

	task := models.ResponseTask{
		ProcessId: command.Process.Pid,
		TaskCommand: command.Path,
		Command: command,
	}

	//logger.log.Println("Process : ", command.Process.Pid)
	//logger.log.Println(command.Stdout)
	return task, nil
}

func IsTerminated(command *exec.Cmd, logger Logger)  bool {
	// not
	return command.ProcessState != nil
}

func Stdout(command *exec.Cmd) string {

	return ioutil.ReadAll(*command.Stdout)
}

func LaunchTestCommands() {
	commands := [] string {"echo start; sleep 10; echo middle", "sleep 20; echo end"}
	status, err := ExecCommandListAsynchronously(commands)

	fmt.Println(status)
	fmt.Println(err)
}
