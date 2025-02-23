package routes

import (
	"github.com/gin-gonic/gin"
	"max-tuts/event-booking-rest-api/models"
	"max-tuts/event-booking-rest-api/utils"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}
	err = user.Authenticate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "Authentication failed"})
		return
	}
	token, err := utils.GenerateJWT(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to generate JWT token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Authentication successful", "user": user, "token": token})
}

func getUsers(context *gin.Context) {
	rows, err := models.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to fetch users"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"users": rows})
}
