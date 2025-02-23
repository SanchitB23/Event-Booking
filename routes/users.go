package routes

import (
	"github.com/gin-gonic/gin"
	"max-tuts/event-booking-rest-api/models"
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
