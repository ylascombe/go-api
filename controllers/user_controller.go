package controllers

import (
	"fmt"
	"net/http"
	"github.com/ylascombe/go-api/services"
	"github.com/ylascombe/go-api/models"
	"github.com/gin-gonic/gin"
)

func FetchAllUsers(c *gin.Context) {
	var users *models.ApiUsers
	var _users []models.TransformedApiUser

	users, err := services.ListApiUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "error message" : err})
		return
	}

	if (len(users.List) <= 0) {
		c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No users found!"})
		return
	}

	//transforms the users
	for _, item := range users.List {
		_users = append(_users, *models.TransformApiUser(item))
	}
	c.JSON(http.StatusOK, _users)
}

func CreateUser(c *gin.Context) {

	var json models.ApiUser

	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status" : http.StatusBadRequest, "message" : "Invalid request.", "error detail": err})
	} else {
		user, err := services.CreateUser(json)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status" : http.StatusInternalServerError, "message" : "Error while creating user", "error detail": err})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"status" : http.StatusCreated,
				"message" : "User created successfully!",
				"user_id": user.ID,
				"Location": fmt.Sprintf("/v1/user/%v", user.ID),
			})
		}
	}


}
