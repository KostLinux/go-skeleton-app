package middleware

import (
	"skeleton-web-app/src/controller"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(headers())

	router.GET("/", controller.GetWelcomeString)
	return router
}
