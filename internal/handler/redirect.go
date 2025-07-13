package handler

import (
	"net/http"

	"github.com/Shobhit150/url_shortner/internal/repository"
	"github.com/Shobhit150/url_shortner/internal/service"
	"github.com/gin-gonic/gin"
)

 func Redirect(c *gin.Context) {
	slug := c.Param("slug")

	longURL, err := service.Redirect(slug)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "URL Not Found"})
	}
	go repository.IncrementClicks(slug)

	c.Redirect(http.StatusFound, longURL)

 }