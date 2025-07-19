package handler

import (
	// "fmt"
	"net/http"
	"time"

	// "github.com/Shobhit150/url_shortner/internal/repository"
	"github.com/Shobhit150/url_shortner/internal/service"
	"github.com/gin-gonic/gin"
)

func RedirectUrl(c *gin.Context) {
	slug := c.Param("slug")
    ip := c.ClientIP()
    userAgent := c.GetHeader("User-Agent")
    referrer := c.GetHeader("Referer")



	
	longURL, expireTime, err := service.Redirect(slug, ip, userAgent, referrer)

	if err != nil {
		// Always check err != nil before err.Error()
		if err.Error() == "URL expired" {
			var expiredAtStr any = nil
			if expireTime != nil {
				expiredAtStr = expireTime.Format(time.RFC3339)
			}
			c.JSON(http.StatusGone, gin.H{"error": "URL expired", "expired_at": expiredAtStr})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "URL Not Found1"})
		return
	}
	// go repository.IncrementClicks(slug)



	c.Redirect(http.StatusFound, longURL)
 }