package routes

import (
	"github.com/gin-gonic/gin"
	"max-tuts/event-booking-rest-api/models"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to register for event"})
		return
	}
}
