package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/Shobhit150/url_shortner/internal/service"
)

func Redirect(c *gin.Context) {
	slug := c.Param("slug")

	url, err := service.Resolve(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Redirect(http.StatusFound, url)
}