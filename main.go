package main

import (
	"github.com/gin-gonic/gin"
	"max-tuts/event-booking-rest-api/db"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	err := server.Run(":8080")
	if err != nil {
		return // exit if the server isn't running
	} // listen and serve on
	println("Server is running on port 8080")
}
