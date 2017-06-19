package controllers

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/models"
	"github.com/gin-gonic/gin"
)


func SSHPublicKeysForEnv(c *gin.Context) {

	envName := c.Param("env-name")

	result, err := services.ListSshPublicKeyForEnv(envName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "error message" : err})
	} else {
		c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : result})
	}
}


func GetEnvironmentAccess(c *gin.Context) {
	envName := c.Param("env-name")

	environmentAccesses, err := services.ListAccessForEnvironment(envName)
	var _environmentAccesses []models.TransformedEnvironmentAccess

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "error message" : err})
		return
	}

	if (environmentAccesses.List == nil) {
		c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : fmt.Sprintf("No access found for environment %s!", envName)})
		return
	}

	//transforms for building a good response
	for i:=0; i<len(environmentAccesses.List); i++ {
		tmp := models.TransformEnvironmentAccess(environmentAccesses.List[i])
		_environmentAccesses = append(_environmentAccesses, *tmp)
	}
	c.JSON(http.StatusOK, _environmentAccesses)
}


func CreateEnvironmentAccess(c *gin.Context) {
	envName := c.Param("env-name")
	userID := c.Param("user-id")
	intUserID, _ := strconv.Atoi(userID)
	uintUserID := uint(intUserID)

	err := services.AddEnvironmentAccess(uintUserID, envName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "message" : "Error while creating environment access", "error detail": err})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Environment access created successfully!", "env_name": envName})
	}
}
