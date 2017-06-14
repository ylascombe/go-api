package utils

import (
	"log"
	"os"
	"fmt"
	"github.com/ylascombe/go-api/config"
	"os/exec"
)

var (
	Log      *log.Logger
)

func NewLog(logpath string) {
	println("LogFile: " + logpath)
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

func info(log string) {
	Log.Println(log)
	if config.LOG_TO_STDOUT {
		fmt.Println(log)
	}
}

func Println(log string, cmd exec.Cmd) {
	Log.Println(log, cmd)
	if config.LOG_TO_STDOUT {
		fmt.Println(log, cmd)
	}
}

//func error(log string, cmd exec.Cmd,err error) {
//	Log.Println(log)
//	fmt.Sprintf("Error when running %s command. Error details: %v\n", cmd, err)
//}