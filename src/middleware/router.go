package middleware

import (
	"skeleton-web-app/src/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(router *gin.Engine) *gin.Engine {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(headers())

	router.GET("/", controller.GetWelcomeString)
	router.HEAD("/", controller.NotBadRequest)
	router.GET("/notbad", controller.NotBadRequest)

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
