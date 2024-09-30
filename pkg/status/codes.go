package status

import (
	"net/http"

	json "skeleton-web-app/pkg/json"

	"github.com/gin-gonic/gin"
)

func OK(ctx *gin.Context, data []byte) {
	var response interface{}
	if err := json.Decode(data, &response); err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func InternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
		"code":    "INTERNAL_SERVER_ERROR",
	})
}

func BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"code":    "BAD_REQUEST",
	})
}

func AbortWithStatusJSON(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
		"code":    http.StatusBadRequest,
	})
}
