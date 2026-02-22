package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hiHandler(c *gin.Context) {
	// Call hello endpoint
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to call hello"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	c.String(200, fmt.Sprintf("Hello says: %s", string(body)))
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hi, eveeryoneee!!!!")
}
