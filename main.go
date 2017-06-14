package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"github.com/ylascombe/go-api/controllers"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	router.HandleFunc("/v1/environment", controllers.Environment)
	router.HandleFunc("/v1/environment/{name}", controllers.Environment)

	router.HandleFunc("/v1/user", controllers.User)

	router.HandleFunc("/v1/environmentAccess/{name}", controllers.EnvironmentAccess)
	router.HandleFunc("/v1/sshKeys/{name}", controllers.SSHPublicKeysForEnv)
	router.HandleFunc("/v1/environmentAccess/{name}/user/{userID}", controllers.EnvironmentAccess)

	router.HandleFunc("/v1/featureTeam", controllers.FeatureTeamCtrl)
	router.HandleFunc("/v1/featureTeam/{name}", controllers.FeatureTeamCtrl)

	router.HandleFunc("/v1/membership/{ftName}", controllers.MembershipCtrl)
	//router.HandleFunc("/manifests", handleListManifests).Methods("GET")

	// XXX keep it at the end of this function
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
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
//			resp := controllers.ApiResponse{ErrorMessage: string(err.Error())}
//			text := utils.Marshall(resp)
//			fmt.Fprintf(writer, text)
//		}
//	} else {
//		writer.WriteHeader(400)
//		resp := controllers.ApiResponse{ErrorMessage: "Given parameters are empty or not valid"}
//		text := utils.Marshall(resp)
//		fmt.Fprintf(writer, text)
//	}
//}
