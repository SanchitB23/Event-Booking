package main

import (
	"github.com/gin-gonic/gin"
	"max-tuts/event-booking-rest-api/db"
	"max-tuts/event-booking-rest-api/models"
	"net/http"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getAllEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080")
	if err != nil {
		return // exit if the server isn't running
	} // listen and serve on
	println("Server is running on port 8080")
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to retrieve events"})
		return
	}
	context.JSON(http.StatusOK, events)
}
