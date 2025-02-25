package routes

import (
	"github.com/gin-gonic/gin"
	"max-tuts/event-booking-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEvent)

	withAuth := server.Group("/")
	withAuth.Use(middlewares.AuthMiddleware)
	withAuth.POST("/events", createEvent)
	withAuth.PUT("/events/:id", updateEvent)
	withAuth.DELETE("/events/:id", deleteEvent)

	withAuth.POST("/events/:id/register", registerForEvent)
	withAuth.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users", getUsers)
}
