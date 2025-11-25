package handler

import (
	"net/http"
	"time"

	"github.com/Shobhit150/url_shortner/internal/service"
	"github.com/gin-gonic/gin"
)


func ShortenURL(c *gin.Context) {
	var body struct {
		URL string `json:"url"`
		CustomSlug string `json:"custom_slug"`
		ExpiresAt  *string `json:"expires_at"`
	}

	if err := c.BindJSON(&body); 
	err != nil || body.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}
	var expiresAt *time.Time
	if body.ExpiresAt != nil && *body.ExpiresAt != "" {
		parsed, err := time.Parse(time.RFC3339, *body.ExpiresAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expires_at format. Use RFC3339."})
			return
		}
		expiresAt = &parsed
	}
	slug, err := service.Shorten(body.URL, body.CustomSlug, expiresAt)
	if err != nil {
		// Custom error handling for known issues
		if err.Error() == "custom slug is already is already taken" {
			c.JSON(http.StatusConflict, gin.H{"error": "Custom slug already exists. Please try another."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL3: " + err.Error()})
		return
	}

	resp := gin.H{
		"shorten_url": "http://localhost:8080/" + slug,
		"slug":        slug,
	}
	if expiresAt != nil {
		resp["expires_at"] = expiresAt.Format(time.RFC3339)
	}
	c.JSON(http.StatusCreated, resp)
}