package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/utils"
	"github.com/ylascombe/go-api/models"
)

func EnvironmentAccess(writer http.ResponseWriter, req *http.Request) {
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
			resp := ApiResponse{ErrorMessage: string(err.Error())}
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
	resp := ApiResponse{ErrorMessage: string(err.Error())}
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
			resp := ApiResponse{ErrorMessage: string(err.Error())}
			text := utils.Marshall(resp)
			fmt.Fprintf(writer, text)
		}
	} else {
		writer.WriteHeader(400)
		resp := ApiResponse{ErrorMessage: "Given parameters are empty or not valid"}
		text := utils.Marshall(resp)
		fmt.Fprintf(writer, text)
	}

}