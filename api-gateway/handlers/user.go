package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	userID := c.Param("id")
	url := fmt.Sprintf("http://127.0.0.1:8081/user/%s", userID)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User service unavailable"})
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}
