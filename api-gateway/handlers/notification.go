package handlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotificationsHandler(c *gin.Context) {
	url := "http://127.0.0.1:8083/notifications"
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notification service unavailable"})
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}
