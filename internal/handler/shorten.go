package handler

import (
	"net/http"

	"github.com/Shobhit150/url_shortner/internal/service"
	"github.com/gin-gonic/gin"
)


func ShortenURL(c *gin.Context) {
	var body struct {
		URL string `json:"url"`
		CustomSlug string `json:"custom_slug"`
	}

	if err := c.BindJSON(&body); 
	err != nil || body.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	slug, err := service.Shorten(body.URL, body.CustomSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}


	c.JSON(http.StatusFound, gin.H{"shorten_url": "http://localhost:8080/" + slug})
}