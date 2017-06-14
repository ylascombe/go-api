package controllers

import (
	"fmt"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/models"
	"github.com/ylascombe/go-api/utils"
	"github.com/gorilla/mux"
)

func MembershipCtrl(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		vars := mux.Vars(req)
		ftName := vars["ftName"]
		members, err := services.ListTeamMembers(ftName)
		genericReadResponse(writer, req, members, err)
	case "POST":
		createMembership(writer, req)
	}
}

func createMembership(writer http.ResponseWriter, request *http.Request) {
	var membership models.Membership
	utils.ReadObjectFromJSONInput(&membership, writer, request)

	if membership.FeatureTeamID != 0 && membership.ApiUserID != 0 {
		_, err := services.CreateMembershipFromIDs(membership.ApiUserID, membership.FeatureTeamID)
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