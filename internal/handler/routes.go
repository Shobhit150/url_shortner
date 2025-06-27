package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	r.POST("/shorten", ShortenURL)
	r.GET("/:slug", Redirect)
}