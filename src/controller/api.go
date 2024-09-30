package controller

import (
	"net/http"
	"skeleton-web-app/pkg/status"

	"github.com/gin-gonic/gin"
)

// GetWelcomeString godoc
// @Summary Get welcome message
// @Description Returns a welcome message
// @Tags welcome
// @Accept  json
// @Produce  json
// @Success 200 {object} WelcomeResponse
// @Failure 400 {object} WelcomeErrorResponse
// @Router / [get]
func GetWelcomeString(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		status.AbortWithStatusJSON(ctx, "Method Not Allowed")
		return
	}
	status.OK(ctx, []byte(`{
        "message": "Welcome to the Skeleton Web App!"
    }`))
}

// NotBadRequest godoc
// @Summary Check for NotBad request
// @Description Returns status 200 if X-Header: NotBad is present
// @Tags check
// @Accept  json
// @Produce  json
// @Success 200 {object} NotBadResponse
// @Failure 400 {object} NotBadErrorResponse
// @Header 200 {string} X-Header "NotBad"
// @Router /notbad [get]
func NotBadRequest(ctx *gin.Context) {
	ctx.Request.Header.Set("X-Header", "VeryBad")
	header := ctx.Request.Header.Get("X-Header")
	if header != "NotBad" {
		status.BadRequest(ctx, `{
            "message": "Invalid header: `+header+`! Only NotBad can be used"
        }`)
		return
	}

	status.OK(ctx, []byte(`{
        "message": "NotBad request received!"
    }`))
}

// Swagger
type WelcomeResponse struct {
	Code    string `json:"code" example:"200"`
	Message string `json:"message" example:"Welcome to the Skeleton Web App!"`
}

type WelcomeErrorResponse struct {
	Code    string `json:"code" example:"400"`
	Message string `json:"message" example:"Method Not Allowed"`
}

type NotBadResponse struct {
	Code    string `json:"code" example:"200"`
	Message string `json:"message" example:"NotBad request received!"`
}

type NotBadErrorResponse struct {
	Code    string `json:"code" example:"400"`
	Message string `json:"message" example:"Invalid header: VeryBad! Only NotBad can be used"`
}
