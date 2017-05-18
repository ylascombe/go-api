package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"sync"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/utils"
)


func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/reactive-platform/target/{target}/manifestversion/{version}", api)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func api(_ http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	target := vars["target"] + ".ini"
	version := vars["version"]
	fmt.Println("target: ", target)
	fmt.Println("version: ", version)

	ansibleCommands := services.BuildCommands(target, version, "fake_path")
	wg := new(sync.WaitGroup)
	wg.Add(3)
	for i:=0; i<len(ansibleCommands); i++ {
		utils.ExecCommand("echo " + ansibleCommands[i],wg)
	}


}
