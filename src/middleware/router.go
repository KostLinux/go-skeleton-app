package middleware

import (
	"skeleton-web-app/src/controller"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(SecurityHeaders())

	router.GET("/", controller.GetWelcomeString)
	return router
}
