package main

import (
	"fmt"
	"html"
	"net/http"
	"github.com/ylascombe/go-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/ylascombe/go-api/database"
)

func main() {

	db := database.NewDBDriver()
	database.AutoMigrateDB(db)

	router := gin.Default()

	users := router.Group("/v1/users")
	{
		users.GET("/", controllers.FetchAllUsers)
		//users.GET("/:name", controllers.GetUser)
		users.POST("/", controllers.CreateUser)
	}

	environments := router.Group("/v1/environments")
	{
		environments.GET("/", controllers.FetchAllEnvironments)
		environments.GET("/:name", controllers.GetEnvironment)
		environments.POST("/:name", controllers.CreateEnvironment)
	}

	// TODO pluralize "environment" term in uri
	// like that : https://github.com/gin-gonic/gin/issues/205
	environmentsAccess := router.Group("/v1/environment/:env-name/access")
	{
		environmentsAccess.GET("/", controllers.GetEnvironmentAccess)
		//environmentsAccess.GET("/:name", controllers.GetEnvironmentAccess)
		environmentsAccess.POST("/:user-id", controllers.CreateEnvironmentAccess)
	}

	environmentsAccessKey := router.Group("/v1/ssh-keys")
	{
		environmentsAccessKey.GET("/:env-name", controllers.SSHPublicKeysForEnv)
	}

	featureTeams := router.Group("/v1/teams")
	{
		featureTeams.GET("/", controllers.FetchAllFeatureTeams)
		//featureTeams.GET("/:name", controllers.GetFeatureTeam)
		//featureTeams.POST("/:name", controllers.CreateFeatureTeam)
		featureTeams.POST("/", controllers.CreateFeatureTeam)
	}

	membership := router.Group("/v1/teams/:team-name/user")
	{
		membership.GET("/", controllers.FetchAllMember)
		//featureTeams.GET("/:name", controllers.GetFeatureTeam)
		//featureTeams.POST("/:name", controllers.CreateFeatureTeam)
		membership.POST("/:user-id", controllers.CreateMembership)
	}

	router.Run(":8090")


	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", Index)
	//
	//router.HandleFunc("/v1/user", controllers.User)
	//
	//router.HandleFunc("/v1/environmentAccess/{name}", controllers.EnvironmentAccess)
	//router.HandleFunc("/v1/sshKeys/{name}", controllers.SSHPublicKeysForEnv)
	//router.HandleFunc("/v1/environmentAccess/{name}/user/{userID}", controllers.EnvironmentAccess)
	//
	//router.HandleFunc("/v1/featureTeam", controllers.FeatureTeamCtrl)
	//router.HandleFunc("/v1/featureTeam/{name}", controllers.FeatureTeamCtrl)
	//
	//router.HandleFunc("/v1/membership/{ftName}", controllers.MembershipCtrl)
	////router.HandleFunc("/manifests", handleListManifests).Methods("GET")
	//
	//// XXX keep it at the end of this function
	//log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

