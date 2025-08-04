package handlers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	url := "http://127.0.0.1:8082/tasks"
	reqBody, _ := io.ReadAll(c.Request.Body)
	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task service unavailable"})
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}
