package middlewares

import (
	"github.com/gin-gonic/gin"
	"max-tuts/event-booking-rest-api/utils"
	"net/http"
)

func AuthMiddleware(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		return
	}
	userId, err := utils.VerifyJWT(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	context.Set("userId", userId)
	context.Next()

}
