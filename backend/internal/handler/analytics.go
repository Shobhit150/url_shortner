package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Shobhit150/url_shortner/internal/repository"
	"github.com/gin-gonic/gin"
)

type AnalyticsResponse struct {
	Slug      string         `json:"slug"`
	ClickCount int           `json:"click_count"`
	Analytics []ClickDetails `json:"analytics"`
}

type ClickDetails struct {
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	Referrer  string    `json:"referrer"`
}

func GetAnalytics(c *gin.Context) {
	slug := c.Param("slug")

	var count int
	err := repository.DB().QueryRow("SELECT count FROM url_clicks WHERE slug = $1", slug).Scan(&count)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Slug not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error: " + err.Error()})
		return
	}

	// Get analytics details (last 20 for example)
	rows, err := repository.DB().Query(
		"SELECT timestamp, ip, user_agent, referrer FROM click_analytics WHERE slug = $1 ORDER BY timestamp DESC LIMIT 20", slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error: " + err.Error()})
		return
	}
	defer rows.Close()

	var analytics []ClickDetails
	for rows.Next() {
		var detail ClickDetails
		if err := rows.Scan(&detail.Timestamp, &detail.IP, &detail.UserAgent, &detail.Referrer); err == nil {
			analytics = append(analytics, detail)
		}
	}

	resp := AnalyticsResponse{
		Slug:      slug,
		ClickCount: count,
		Analytics: analytics,
	}

	c.JSON(http.StatusOK, resp)
}