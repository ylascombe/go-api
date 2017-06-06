package utils

import (
	"log"
	"os"
	"fmt"
	"github.com/ylascombe/go-api/config"
	"os/exec"
)

type Logger struct {
	log *log.Logger
}

func NewLog(logpath string) Logger{
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}
	return Logger{log.New(file, "", log.LstdFlags|log.Lshortfile)}
}

func (logger Logger) Info(log string) {
	logger.log.Println(log)
	if config.LOG_TO_STDOUT {
		fmt.Println(log)
	}
}

func (logger Logger) Println(log string, cmd exec.Cmd) {
	logger.log.Println(log, cmd)
	if config.LOG_TO_STDOUT {
		fmt.Println(log, cmd)
	}
}

func (logger Logger) Error(log string, cmd exec.Cmd,err error) {
	logger.log.Println(log)
}
