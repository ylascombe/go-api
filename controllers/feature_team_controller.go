package controllers

import (
	"fmt"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/models"
	"github.com/gin-gonic/gin"
)

func FetchAllFeatureTeams(c *gin.Context) {
	var teams *models.FeatureTeams
	var _teams []models.TransformedFeatureTeam

	teams, err := services.ListFeatureTeams()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "error message" : err})
		return
	}

	if (len(teams.List) <= 0) {
		c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No feature team found!"})
		return
	}

	//transforms the features teams
	for _, item := range teams.List {
		transformed := models.TransformFeatureTeam(item)
		_teams = append(_teams, *transformed)
	}
	c.JSON(http.StatusOK, _teams)
}

func CreateFeatureTeam(c *gin.Context) {

	var json models.FeatureTeam

	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status" : http.StatusBadRequest, "message" : "Invalid request.", "error detail": err})
	} else {
		res, err := services.CreateFeatureTeam(json)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "message" : "Error while creating user", "error detail": err})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"status" : http.StatusCreated,
				"message" : "User created successfully!",
				"Location": fmt.Sprintf("/v1/teams/%v", res.ID),
				"featureteam_id": res.ID,
			})
		}
	}


}
