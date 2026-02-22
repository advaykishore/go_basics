package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

// In-memory inventory
var inventory = map[string]int{
	"mouse":    100,
	"keyboard": 50,
}

func orderHandler(c *gin.Context) {
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Just return order info (NO inventory update)
	c.JSON(200, req)
}

func updateInventoryHandler(c *gin.Context) {
	var input OrderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Convert request → JSON
	data, _ := json.Marshal(input)

	resp, err := http.Post("http://localhost:8080/order", "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.JSON(500, gin.H{"error": "order service failed"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Parse order response
	var orderResp OrderRequest
	json.Unmarshal(body, &orderResp)

	// Reduce inventory
	if inventory[orderResp.Product] < orderResp.Quantity {
		c.JSON(400, gin.H{"error": "not enough stock"})
		return
	}

	inventory[orderResp.Product] -= orderResp.Quantity

	c.JSON(200, gin.H{
		"message":   "inventory updated",
		"product":   orderResp.Product,
		"remaining": inventory[orderResp.Product],
	})
}

func inventoryHandler(c *gin.Context) {
	c.JSON(200, inventory)
}
