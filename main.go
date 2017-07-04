package main

import (
	"arc-api/controllers"
	"arc-api/database"
	"github.com/itsjamie/gin-cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	db := database.NewDBDriver()
	database.AutoMigrateDB(db)

	router := gin.New()
	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))


	users := router.Group("/v1/users")
	{
		users.GET("/", controllers.FetchAllUsers)
		//users.GET("/:name", controllers.GetUser)
		users.POST("/", controllers.CreateUser)
	}

	environments := router.Group("/v1/environments")
	{
		environments.GET("/", controllers.FetchAllEnvironments)
		environments.GET("/:env-name", controllers.GetEnvironment)
		environments.POST("/:env-name", controllers.CreateEnvironment)
	}

	environmentsAccess := router.Group("/v1/environments/:env-name/access")
	{
		environmentsAccess.GET("/", controllers.GetEnvironmentAccess )
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
}
