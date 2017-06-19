package controllers

import (
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/models"
	"github.com/gin-gonic/gin"
)


func FetchAllEnvironments(c *gin.Context) {
	var environments *[]models.Environment
	var _environments []models.TransformedEnvironment

	environments, err := services.ListEnvironment()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "error message" : err})
		return
	}

	if (len(*environments) <= 0) {
		c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No environment found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range *environments {
		_environments = append(_environments, models.TransformedEnvironment{ID: item.ID, Name: item.Name})
	}
	//c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _environments})
	c.JSON(http.StatusOK, _environments)
}

func GetEnvironment(c *gin.Context) {
	var environment *models.Environment
	envName := c.Param("name")

	environment, err := services.GetEnvironmentByName(envName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "data" : ""})
		return
	}

	if (environment.ID == 0) {
		c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No environment found!"})
		return
	}

	_environment := models.TransformedEnvironment{ID: environment.ID, Name: environment.Name}
	c.JSON(http.StatusOK, _environment)
}

func CreateEnvironment(c *gin.Context) {
	envName := c.Param("name")
	_, err := services.CreateEnvironment(envName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "message" : "Error while creating environment", "error detail": err})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Environment created successfully!", "env_name": envName})
	}
}
