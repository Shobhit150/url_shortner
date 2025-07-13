package handler

import (
	"net/http"

	"github.com/Shobhit150/url_shortner/internal/service"
	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
	slug := c.Param("slug")
	clicks, err := service.GetClicks(slug)
	if (err != nil) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"slug": slug, "clicks": clicks})
}