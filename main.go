package main

import (
	"github.com/gin-gonic/gin"
	"max-tuts/event-booking-rest-api/db"
	"max-tuts/event-booking-rest-api/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		println("Error starting server:", err)
		return // exit if the server isn't running
	} // listen and serve on

	println("Server is running on port 8080")
}
