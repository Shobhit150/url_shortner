package handler

import (
	"github.com/gin-gonic/gin"
    "net/http"
    "github.com/Shobhit150/url_shortner/internal/service"
)

func ShortenURL(c *gin.Context) {
	var body struct {
		URL string `json:"url"`
	}
	if err:= c.BindJSON(&body); err != nil || body.URL =="" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	slug, err := service.Shorten(body.URL) 
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"shorten_url" : "https://localhost:8080" + slug})

}