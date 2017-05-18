package utils

import (
	"sync"
	"fmt"
	"os/exec"
	"errors"
	"github.com/ylascombe/go-api/models"
	"flag"
)

var (
	logpath = flag.String("logpath", "/tmp/go-api-log.log", "API logs")
	commands = make(map[models.ResponseTask]*exec.Cmd)
)

func ExecCommand(cmd string, wg *sync.WaitGroup) (int, error) {
	fmt.Println("Prepare to execute command : ", cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		err_msg := fmt.Sprintf("Error when running %s command. Error details: %v\n", cmd, err)
		return -1, errors.New(err_msg)
	}
	fmt.Printf("\tCommand result : \n\t%s", out)
	wg.Done()
	return 0, nil
}

func ExecCommandsAsynchronously(cmd []string) (models.ResponseTask, error) {

	fmt.Println("start")
	for i:=0;i<len(cmd);i++ {

		command, _ := ExecCommandAsynchronously(cmd[i])

		commands[command].Wait()
		Log.Println("La tache " + cmd[i] + " a terminé, passage à la suivante")

	}
	fmt.Println("end")

	return models.ResponseTask{}, nil
}

func ExecCommandAsynchronously(cmd string) (models.ResponseTask, error) {

	fmt.Println("Prepare to execute command : ", cmd)
	command := exec.Command("sh", "-c", cmd)

	flag.Parse()
	NewLog(*logpath)
	Log.Println(cmd)
	err := command.Start()
	if err != nil {
		err_msg := fmt.Sprintf("Error when running %s command. Error details: %v\n", cmd, err)
		return models.ResponseTask{}, errors.New(err_msg)
	}

	task := models.ResponseTask{
		ProcessId: command.Process.Pid,
		TaskCommand: cmd,
	}

	commands[task] = command
	Log.Println("Process : ", command.Process.Pid)
	return task, nil
}

func IsTerminated(processId int) bool {

	return true

}

func LaunchTestCommands() {
	commands := [] string {"echo start; sleep 10; echo middle", "sleep 20; echo end"}
	status, err := ExecCommandsAsynchronously(commands)
	Log.Println(status)
	Log.Println(err)
}