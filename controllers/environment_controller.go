package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/utils"
)

func Environment(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		listEnvironments(writer, req)
	case "POST":
		createEnvironment(writer, req)
	}
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
	resp := ApiResponse{ErrorMessage: string(err.Error())}
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
		resp := ApiResponse{ErrorMessage: string(err.Error())}
		text := utils.Marshall(resp)
		fmt.Fprintf(writer, text)
	}
}