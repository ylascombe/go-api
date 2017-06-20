package controllers

import (
	"fmt"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

func CreateMembership(c *gin.Context) {
	teamName := c.Param("team-name")
	userID := c.Param("user-id")
	intUserID, _ := strconv.Atoi(userID)
	uintUserID := uint(intUserID)

	featureTeam, err := services.GetFeatureTeamFromName(teamName)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : fmt.Sprintf("No feature team name %s found!", teamName)})
		return
	}

	_, err = services.CreateMembershipFromIDs(uintUserID, featureTeam.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "message" : "Error while creating environment access", "error detail": err})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : fmt.Sprintf("User %s has been added to team %s!", userID, teamName)})
		// TODO return location
	}
}

func FetchAllMember(c *gin.Context) {
	teamName := c.Param("team-name")
	var _memberships []models.TransformedMembership

	memberships, err := services.ListTeamMembers(teamName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "error message" : err})
		return
	}

	if (len(memberships.List) <= 0) {
		// choice : if no membership found, return a HTTP status code 200 with an empty array
		_memberships = make([]models.TransformedMembership, 0)
	}

	//transforms features teams
	for _, item := range memberships.List {
		tmp := models.TransformMembership(item)
		_memberships = append(_memberships, *tmp)
	}
	c.JSON(http.StatusOK, _memberships)
}
