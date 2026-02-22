package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hi", hiHandler)
	r.GET("/hello", helloHandler)
	r.POST("/order", orderHandler)
	r.POST("/update_inventory", updateInventoryHandler)
	r.GET("/inventory", inventoryHandler)

	r.Run(":8080") // start server
}
