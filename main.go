package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ylascombe/go-api/config"
	"github.com/ylascombe/go-api/models"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/utils"
	"html"
	"log"
	"net/http"
	"strconv"
)

type apiResponse struct {
	ErrorMessage string `yaml:"error,omitempty"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/reactive-platform/target/{target}/manifestversion/{version}", api)

	router.HandleFunc("/v1/environment", environment)
	router.HandleFunc("/v1/environment/{name}", environment)

	router.HandleFunc("/v1/user", user)

	router.HandleFunc("/v1/environmentAccess/{name}", environmentAccess)
	router.HandleFunc("/v1/environmentAccess/{name}/user/{userID}", environmentAccess)

	//router.HandleFunc("/manifests", handleListManifests).Methods("GET")

	// TODO remove me when tests done
	router.HandleFunc("/testCommands", launchCommand)

	// XXX keep it at the end of this function
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

func listEnvironments(writer http.ResponseWriter, r *http.Request) {
	envs, err := services.ListEnvironment()

	if err == nil {
		bytes, err := json.Marshal(*envs)

		if err == nil {
			fmt.Fprintf(writer, string(bytes))

			if err == nil {
				// YAML Marshalling : text := utils.Marshall(envs)
				text := utils.Marshall(envs)

				writer.WriteHeader(200)
				fmt.Fprintf(writer, string(text))
				return
			}
		}
	}


	// if this code is executed, so there is an error
	writer.WriteHeader(500)
	resp := apiResponse{ErrorMessage: string(err.Error())}
	text := utils.Marshall(resp)
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
		listEnvironments(writer, req)
	case "POST":
		createEnvironment(writer, req)
	}
}

func environmentAccess(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	fmt.Println("Method : ", req.Method)
	switch req.Method {
	case "GET":
		listEnvironmentAccesses(writer, req, name)

	case "PUT":
		userID := vars["userID"]
		intUserID, _ := strconv.Atoi(userID)
		uintUserID := uint(intUserID)
		fmt.Println("environmentAccess2")

		err := services.AddEnvironmentAccess(uintUserID, name)
		fmt.Println("environmentAccess3")
		fmt.Println(err)

		if err == nil {
			fmt.Println("environmentAccess - 200")
			writer.WriteHeader(200)
			return
		} else {
			fmt.Println("environmentAccess - 409")
			writer.WriteHeader(409)
			resp := apiResponse{ErrorMessage: string(err.Error())}
			text := utils.Marshall(resp)
			fmt.Fprintf(writer, text)
		}
	default:
		fmt.Println("environmentAccess - 404")
		writer.WriteHeader(404)
	}
}

func listEnvironmentAccesses(writer http.ResponseWriter, request *http.Request, name string) {
	environmentAccesses, err := services.ListAccessForEnvironment(name)

	fmt.Println("name", name)
	fmt.Println(environmentAccesses)
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

func createEnvironmentAccess(writer http.ResponseWriter, request *http.Request) {
	var envAccess models.EnvironmentAccess

	//readObjectAndCallCreateFunction(envAccess, services.CreateEnvironmentAccess, writer, request)
	utils.ReadObjectFromJSONInput(&envAccess, writer, request)

	if envAccess.IsValid() {
		_, err := services.CreateEnvironmentAccess(&envAccess)
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
