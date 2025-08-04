package handlers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	url := "http://127.0.0.1:8081/login"
	reqBody, _ := io.ReadAll(c.Request.Body)
	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request"})
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
