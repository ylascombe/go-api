package utils

import (
	"sync"
	"fmt"
	"os/exec"
	"errors"
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