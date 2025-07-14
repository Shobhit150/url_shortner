package handler

import (
	"fmt"
	"net/http"

	"github.com/Shobhit150/url_shortner/internal/repository"
	"github.com/Shobhit150/url_shortner/internal/service"
	"github.com/gin-gonic/gin"
)

 func RedirectUrl(c *gin.Context) {
	slug := c.Param("slug")

	fmt.Println("redirect file slug : ", slug)
	fmt.Println("redirect file c : ", c)
	longURL, err := service.Redirect(slug)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "URL Not Found1"})
		return
	}
	go repository.IncrementClicks(slug)

	c.Redirect(http.StatusFound, longURL)
 }