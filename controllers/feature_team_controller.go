package controllers

import (
	"fmt"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/models"
	"github.com/ylascombe/go-api/utils"
)

func FeatureTeamCtrl(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		fts, err := services.ListFeatureTeams()
		genericReadResponse(writer, req, fts, err)
	case "POST":
		createFeatureTeam(writer, req)
	}
}

func createFeatureTeam(writer http.ResponseWriter, request *http.Request) {
	var ft models.FeatureTeam
	utils.ReadObjectFromJSONInput(&ft, writer, request)

	if ft.IsValid() {
		_, err := services.CreateFeatureTeam(ft)
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