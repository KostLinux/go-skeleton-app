package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotBadRequest checks if the "NotBad" header is set to "true" and responds accordingly
func GetWelcomeString(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the API!",
	})
}
