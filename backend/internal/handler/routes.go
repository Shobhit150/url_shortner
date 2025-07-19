 package handler

 import (
	"github.com/gin-gonic/gin"
 )

 func RegisterRouters(r *gin.Engine) {
	r.POST("/shorten", ShortenURL)
	r.GET("/stats/:slug", GetStats)
	r.GET("/analytics/:slug", GetAnalytics)
	r.GET("/:slug", RedirectUrl)
	
 }