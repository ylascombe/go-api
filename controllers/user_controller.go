package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/utils"
	"github.com/ylascombe/go-api/models"
)

func User(writer http.ResponseWriter, req *http.Request) {
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
	resp := ApiResponse{ErrorMessage: string(err.Error())}
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