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

type apiResponse struct {
	ErrorMessage string      `yaml:"error,omitempty"`
//	ID           string      `json:"id,omitempty"`
//	Result       interface{} `json:"result,omitempty"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/reactive-platform/target/{target}/manifestversion/{version}", api)
	router.HandleFunc("/testCommands", launchCommand)

	router.HandleFunc("/v1/environment", environment)
	router.HandleFunc("/v1/environment/{name}", environment)
	//router.HandleFunc("/manifests", handleListManifests).Methods("GET")
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

	logger := utils.NewLog("/tmp/api.txt")
	go utils.ExecCommandListAsynchronously(ansibleCommands, logger)
	fmt.Fprintf(writer, "Les commandes ont été lancées")
}

func launchCommand(w http.ResponseWriter, r *http.Request) {
	go utils.LaunchTestCommands()
	fmt.Fprintf(w, "Les commandes ont été lancées, %q", html.EscapeString(r.URL.Path))
}

func isTerminated(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Statut des commandes, %q", html.EscapeString(r.URL.Path))
}

func environment(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		listEnvironments(writer, req)
	case "POST":
		createEnvironment(writer, req)
	}
}

func listEnvironments(writer http.ResponseWriter, r *http.Request) {
	envs := services.ListEnvironment()
	text := utils.Marshall(*envs)
	fmt.Fprintf(writer, text)
}

func createEnvironment(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	name := vars["name"]

	_, err := services.CreateEnvironment(name)

	if err == nil {
		writer.WriteHeader(200)
	} else {
		writer.WriteHeader(409)
		resp := apiResponse{ErrorMessage: string(err.Error())}
		text := utils.Marshall(resp)
		fmt.Fprintf(writer, text)

	}
}
