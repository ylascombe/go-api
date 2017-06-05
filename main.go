package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/utils"
	"github.com/ylascombe/go-api/config"
)


func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/reactive-platform/target/{target}/manifestversion/{version}", api)
	router.HandleFunc("/testCommands", launchCommand)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func api(writer http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	target := vars["target"] + ".ini"
	version := vars["version"]
	fmt.Println("target: ", target)
	fmt.Println("target: ", target)

	fmt.Println("version: ", version)

	ansibleCommands := services.BuildCommands(target, version, config.VAULT_SECRET_FILE)

	go utils.ExecCommandListAsynchronously(ansibleCommands)
	fmt.Fprintf(writer, "Les commandes ont été lancées")
}

func launchCommand(w http.ResponseWriter, r *http.Request) {
	go utils.LaunchTestCommands()
	fmt.Fprintf(w, "Les commandes ont été lancées, %q", html.EscapeString(r.URL.Path))
}

func isTerminated(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Statut des commandes, %q", html.EscapeString(r.URL.Path))
}


