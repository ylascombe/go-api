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
	"encoding/json"
	"github.com/ylascombe/go-api/models"
	"strconv"
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
	router.HandleFunc("/v1/user", user)
	router.HandleFunc("/v1/environmentAccess/{name}", environmentAccess)
	router.HandleFunc("/v1/environmentAccess/{name}/user/{userID}", environmentAccess)
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

func environmentAccess(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	switch req.Method {
	case "GET":
		services.ListAccessForEnvironment(name)
	case "PUT":
		userID := vars["userID"]
		intUserID, _ := strconv.Atoi(userID)
		uintUserID := uint(intUserID)
		fmt.Println("environmentAccess")

		err := services.AddEnvironmentAccess(uintUserID, name)
		fmt.Println("environmentAccess2")

		if err == nil {
			writer.WriteHeader(200)
			return
		} else {
			writer.WriteHeader(409)
			resp := apiResponse{ErrorMessage: string(err.Error())}
			text := utils.Marshall(resp)
			fmt.Fprintf(writer, text)
		}
	}
}

func listEnvironments(writer http.ResponseWriter, r *http.Request) {
	envs, err := services.ListEnvironment()
	text := utils.Marshall(*envs)
	fmt.Fprintf(writer, text)

	if err == nil {
		// YAML Marshalling : text := utils.Marshall(envs)
		text, err := json.Marshal(envs)

		if err == nil {
			writer.WriteHeader(200)
			fmt.Fprintf(writer, string(text))
			return
		}
	}

	// if this code is executed, so there is an error
	writer.WriteHeader(500)
	resp := apiResponse{ErrorMessage: string(err.Error())}
	text = utils.Marshall(resp)
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

func user(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		listUsers(writer, req)
	case "POST":
		createUser(writer, req)
	}
}

func listUsers(writer http.ResponseWriter, r *http.Request) {
	users, err := services.ListApiUser()

	if err == nil {
		// YAML Marshalling : text := utils.Marshall(users)
		text, err := json.Marshal(users)

		if err == nil {
			writer.WriteHeader(200)
			fmt.Fprintf(writer, string(text))
			return
		}
	}

	// if this code is executed, so there is an error
	writer.WriteHeader(500)
	resp := apiResponse{ErrorMessage: string(err.Error())}
	text := utils.Marshall(resp)
	fmt.Fprintf(writer, text)

}

func createUser(writer http.ResponseWriter, request *http.Request) {
	var user models.ApiUser
	utils.ReadObjectFromJSONInput(&user, writer, request)

	if user.IsValid() {
		_, err := services.CreateUser(user)
		if err == nil {
			writer.WriteHeader(200)
		} else {
			writer.WriteHeader(409)
			resp := apiResponse{ErrorMessage: string(err.Error())}
			text := utils.Marshall(resp)
			fmt.Fprintf(writer, text)
		}
	} else {
		writer.WriteHeader(400)
		resp := apiResponse{ErrorMessage: "Given parameters are empty or not valid"}
		text := utils.Marshall(resp)
		fmt.Fprintf(writer, text)
	}
}


func environment(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		listEnvAccess(writer, req)
	case "POST":
		createEnvAccess(writer, req)
	}
}
func listEnvAccess(writer http.ResponseWriter, request *http.Request) {
	environmentAccesses, err := services.ListEnvironmentAccesses()

	if err == nil {
		// for YAML Marshalling : text := utils.Marshall(users)
		text, err := json.Marshal(environmentAccesses)

		if err == nil {
			writer.WriteHeader(200)
			fmt.Fprintf(writer, string(text))
			return
		}
	}

	// if this code is executed, so there is an error
	writer.WriteHeader(500)
	resp := apiResponse{ErrorMessage: string(err.Error())}
	text := utils.Marshall(resp)
	fmt.Fprintf(writer, text)
}

func createEnvAccess(writer http.ResponseWriter, request *http.Request) {
	var envAccess models.EnvironmentAccess

	//readObjectAndCallCreateFunction(envAccess, services.CreateEnvironmentAccess, writer, request)
	utils.ReadObjectFromJSONInput(&envAccess, writer, request)

	if envAccess.IsValid() {
		_, err := services.CreateEnvironmentAccess(envAccess)
		if err == nil {
			writer.WriteHeader(200)
		} else {
			writer.WriteHeader(409)
			resp := apiResponse{ErrorMessage: string(err.Error())}
			text := utils.Marshall(resp)
			fmt.Fprintf(writer, text)
		}
	} else {
		writer.WriteHeader(400)
		resp := apiResponse{ErrorMessage: "Given parameters are empty or not valid"}
		text := utils.Marshall(resp)
		fmt.Fprintf(writer, text)
	}

}

// TODO comprendre pourquoi cette factorisation ne marche pas (c'est problème l'interface Validable)
//func readObjectAndCallCreateFunction(toCreate models.Validable, fct func(models.Validable) (models.Validable, error), writer http.ResponseWriter, request *http.Request) {
//
//	utils.ReadObjectFromJSONInput(&toCreate, writer, request)
//
//	if toCreate.IsValid() {
//		_, err := fct(toCreate)
//		if err == nil {
//			writer.WriteHeader(200)
//		} else {
//			writer.WriteHeader(409)
//			resp := apiResponse{ErrorMessage: string(err.Error())}
//			text := utils.Marshall(resp)
//			fmt.Fprintf(writer, text)
//		}
//	} else {
//		writer.WriteHeader(400)
//		resp := apiResponse{ErrorMessage: "Given parameters are empty or not valid"}
//		text := utils.Marshall(resp)
//		fmt.Fprintf(writer, text)
//	}
//}
